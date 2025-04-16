package metrics

import (
	"os/exec"
	"strings"
)

func UpdateSSHConnections() {
	out, err := exec.Command("who").Output()
	if err != nil {
		return
	}

	SSHConnectionsByUser.Reset()
	count := 0
	lines := strings.Split(string(out), "\n")
	userCounts := make(map[string]int)

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		user := fields[0]
		if strings.Contains(line, "pts/") {
			count++
			userCounts[user]++
		}
	}

	SSHConnections.Set(float64(count))
	for user, c := range userCounts {
		SSHConnectionsByUser.WithLabelValues(user).Set(float64(c))
	}
}
