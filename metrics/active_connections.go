package metrics

import (
	"os/exec"
	"strings"
)

func UpdateSSHConnections() {
	// Explicitly point to the utmp file
	out, err := exec.Command("who", "/var/run/utmp").Output()
	if err != nil {
		return
	}

	SSHConnectionsByUser.Reset()

	totalCount := 0
	userCounts := make(map[string]int)

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

	SSHConnections.Set(float64(totalCount))
	for user, c := range userCounts {
		SSHConnectionsByUser.WithLabelValues(user).Set(float64(c))
	}
}
