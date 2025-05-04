package metrics

import (
	"os/exec"
	"strings"
	"time"
)

// UpdateLoginsToday sets SSHLoginsToday based on unique SSH sessions today
func UpdateLoginsToday() {
	out, err := exec.Command("last").Output()
	if err != nil {
		return
	}
	lines := strings.Split(string(out), "\n")
	today := time.Now().Format("Jan _2")
	uniqueSessions := make(map[string]bool)

	for _, line := range lines {
		if strings.Contains(line, "pts/") && strings.Contains(line, today) {
			// Use entire line as unique key or parts of it
			sessionKey := strings.Join(strings.Fields(line)[:5], " ")
			uniqueSessions[sessionKey] = true
		}
	}

	SSHLoginsToday.Set(float64(len(uniqueSessions)))
}
