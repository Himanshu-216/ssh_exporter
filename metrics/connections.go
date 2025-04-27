package metrics

import (
	// "fmt"
	// "log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Update the SSH connection counts based on the logs and fingerprints
func UpdateSSHConnections() {
	out, err := exec.Command("who").Output()
	if err != nil {
		return
	}
	// Reset metrics at the beginning
	SSHConnectionsByUser.Reset()
	SSHConnectionsByFingerprint.Reset()

	totalCount := 0
	userCounts := make(map[string]int)
	fingerprintCounts := make(map[string]int)

	// Map: fingerprint => comment (usually the actual user's name/email at the end of the pub key)
	fingerprintLabelMap := getAllUserAuthorizedKeyFingerprints()


	// Parsing "who" output for active sessions
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 2 || !strings.Contains(line, "pts/") {
			continue
		}

		username := fields[0]
		totalCount++
		userCounts[username]++

		// Get the fingerprint from SSH logs
		fp := getFingerprintFromAuthLog(username)
		if fp != "" {
			// Map fingerprint to comment (i.e., user name or other useful info)
			label := fp
			if comment, ok := fingerprintLabelMap[fp]; ok {
				label = comment // Use the comment if fingerprint matches
			}

			// Increment connection count for this fingerprint (based on comment)
			fingerprintCounts[label]++
		}
	}
	

	// Update the total SSH connections
	SSHConnections.Set(float64(totalCount))
	for user, c := range userCounts {
		SSHConnectionsByUser.WithLabelValues(user).Set(float64(c))
	}
	for label, c := range fingerprintCounts {
		SSHConnectionsByFingerprint.WithLabelValues(label).Set(float64(c))
		print(SSHConnectionsByFingerprint)
	}
}


// Parses all authorized_keys files and returns map[fingerprint] = key_comment
func getAllUserAuthorizedKeyFingerprints() map[string]string {
	matches := make(map[string]string)

	// Read /etc/passwd to get the home directories of users
	data, err := os.ReadFile("/etc/passwd")
	if err != nil {
		return matches
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) < 6 {
			continue
		}
		homeDir := parts[5]
		authKeys := filepath.Join(homeDir, ".ssh", "authorized_keys")
		content, err := os.ReadFile(authKeys)
		if err != nil {
			continue
		}

		// Loop through each key in the authorized_keys file
		for _, keyLine := range strings.Split(string(content), "\n") {
			if strings.HasPrefix(keyLine, "ssh-") {
				fp, comment := fingerprintOfKey(keyLine)
				if fp != "" {
					matches[fp] = comment // Store fingerprint => comment mapping
				}
			}
		}
	}
	return matches
}

func fingerprintOfKey(keyLine string) (string, string) {
	cmd := exec.Command("ssh-keygen", "-lf", "/dev/stdin")
	cmd.Stdin = strings.NewReader(keyLine)
	out, err := cmd.Output()
	if err != nil {
		return "", ""
	}
	fields := strings.Fields(string(out))
	if len(fields) >= 3 {
		fingerprint := fields[1]        // e.g., SHA256:xxxxx
		comment := strings.Join(fields[2:], " ") // In case comment has spaces
		return fingerprint, comment
	}
	return "", ""
}


// Parses recent SSHD logs to find fingerprint used by a user
func getFingerprintFromAuthLog(username string) string {
	logOutput, err := exec.Command("journalctl", "_COMM=sshd", "--since=1 hour ago", "-g", "Accepted publickey").Output()
	if err != nil {
		return ""
	}
	lines := strings.Split(string(logOutput), "\n")
	for _, line := range lines {
		if strings.Contains(line, username) && strings.Contains(line, "SHA256:") {
			parts := strings.Split(line, "SHA256:")
			if len(parts) > 1 {
				return "SHA256:" + strings.Fields(parts[1])[0]
			}
		}
	}
	return ""
}

