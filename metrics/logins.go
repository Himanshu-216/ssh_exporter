package metrics

import (
	"os/exec"
	"strings"
	"time"
)

func UpdateLoginsToday() {
	out, err := exec.Command("last").Output()
	if err != nil {
		return
	}
	lines := strings.Split(string(out), "\n")
	today := time.Now().Format("Jan 2")
	count := 0

	for _, line := range lines {
		// pts/ indicates a pseudo-terminal, typically used by SSH sessions
		if strings.Contains(line, "pts/") && strings.Contains(line, today) {
			count++
		}
	}

	SSHLoginsToday.Set(float64(count))
}
