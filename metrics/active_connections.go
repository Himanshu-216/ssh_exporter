package metrics

import (
	// "os"
	"os/exec"
	"strings"
	// "fmt"
)

// Update the SSH connection counts based on the "who" output
func UpdateSSHConnections() {
	out, err := exec.Command("who").Output()
	if err != nil {
		return
	}
	// Reset metrics at the beginning
	SSHConnectionsByUser.Reset()

	totalCount := 0
	userCounts := make(map[string]int)

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
	}

	// Update the total SSH connections
	SSHConnections.Set(float64(totalCount))
	for user, c := range userCounts {
		SSHConnectionsByUser.WithLabelValues(user).Set(float64(c))
	}
}
