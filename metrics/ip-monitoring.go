package metrics

import (
	"os/exec"
	"regexp"
	"strings"
	"time"
)

// UpdateLoginMetrics updates the metrics for successful and failed logins.
func LoginsmonitorwithIP() {
	// Get failed login attempts (from lastb)
	failedOut, err := exec.Command("lastb").Output()
	if err != nil {
		return
	}
	failedLines := strings.Split(string(failedOut), "\n")
	today := time.Now().Format("Jan 2")
	failedCount := 0
	failedLoginIPCount := make(map[string]int)

	// Regex to extract IP addresses from the output
	ipRegex := regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`)

	// Process failed logins
	for _, line := range failedLines {
		if strings.Contains(line, today) {
			failedCount++
			ipMatch := ipRegex.FindString(line)
			if ipMatch != "" {
				failedLoginIPCount[ipMatch]++
			}
		}
	}

	// Set total failed logins for today
	SSHFailedLoginsToday.Set(float64(failedCount))

	// Update Prometheus gauge for each failed login IP
	for ip, ipCount := range failedLoginIPCount {
		SSHFailedLoginsByIP.WithLabelValues(ip).Set(float64(ipCount))
	}

	// Get successful login attempts (from last)
	successOut, err := exec.Command("last").Output()
	if err != nil {
		return
	}
	successLines := strings.Split(string(successOut), "\n")
	successCount := 0
	successfulLoginIPCount := make(map[string]int)

	// Process successful logins
	for _, line := range successLines {
		if strings.Contains(line, today) {
			successCount++
			ipMatch := ipRegex.FindString(line)
			if ipMatch != "" {
				successfulLoginIPCount[ipMatch]++
			}
		}
	}

	// Set total successful logins for today
	SSHLoginsToday.Set(float64(successCount))

	// Update Prometheus gauge for each successful login IP
	for ip, ipCount := range successfulLoginIPCount {
		SSHSuccessfulLoginsByIP.WithLabelValues(ip).Set(float64(ipCount))
	}
}
