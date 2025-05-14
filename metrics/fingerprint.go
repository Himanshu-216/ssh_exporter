package metrics

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

// Session struct to track session information
type Session struct {
	StartTime      time.Time
	User           string
	RsaFingerprint string
	SessionID      string
}

// monitorAuthLog reads the auth log and updates metrics based on login/logout events
func monitorAuthLog(logFilePath string) error {
	file, err := os.Open(logFilePath)
	if err != nil {
		return fmt.Errorf("error opening log file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Define regex patterns for login, logout, and other session events
	loginRegex := regexp.MustCompile(`sshd\[\d+\]: Accepted publickey for (\w+) from .*\s+ssh-rsa (\S+)`)
	sessionStartRegex := regexp.MustCompile(`systemd-logind\[\d+\]: New session (\d+) of user (\w+)`)
	sessionEndRegex := regexp.MustCompile(`systemd-logind\[\d+\]: Removed session (\d+)`)

	for scanner.Scan() {
		line := scanner.Text()

		// SSH login event: Capture user and RSA fingerprint
		if loginMatch := loginRegex.FindStringSubmatch(line); loginMatch != nil {
			user := loginMatch[1]
			rsaFingerprint := loginMatch[2]
			sessionID := generateSessionID(user, rsaFingerprint)

			// Store session start time
			sessions[sessionID] = &Session{
				StartTime:      time.Now(),
				User:           user,
				RsaFingerprint: rsaFingerprint,
				SessionID:      sessionID,
			}

			// Optionally log for debugging
			fmt.Printf("SSH login detected: User %s, RSA Fingerprint: %s, SessionID: %s\n", user, rsaFingerprint, sessionID)
		}

		// Systemd session start event: New session is created
		if startMatch := sessionStartRegex.FindStringSubmatch(line); startMatch != nil {
			sessionID := startMatch[1]
			user := startMatch[2]

			// Optionally log for debugging
			fmt.Printf("Session started: SessionID: %s, User: %s\n", sessionID, user)
		}

		// Session end event: Session removed (logout)
		if endMatch := sessionEndRegex.FindStringSubmatch(line); endMatch != nil {
			sessionID := endMatch[1]

			// Check if this session exists
			session, exists := sessions[sessionID]
			if exists {
				// Calculate session duration
				duration := time.Since(session.StartTime).Seconds()

				// Update the metric with the session duration
				SSHSessionDuration.WithLabelValues(sessionID, session.User, session.RsaFingerprint).Set(duration)

				// Optionally log for debugging
				fmt.Printf("Session ended: SessionID: %s, User: %s, Duration: %.2f seconds\n", sessionID, session.User, duration)

				// Clean up the session map
				delete(sessions, sessionID)
			}
		}
	}

	// Handle any scanner errors
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading log file: %v", err)
	}

	return nil
}

// Helper function to generate a unique session ID (user + RSA fingerprint)
func generateSessionID(user, rsaFingerprint string) string {
	return fmt.Sprintf("%s-%s", user, rsaFingerprint)
}
