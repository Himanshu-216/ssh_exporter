package metrics

import (
	"os/exec"
	"strings"
	"time"
)

func UpdateFailedLogins() {
	out, err := exec.Command("lastb").Output()
	if err != nil {
		return
	}
	lines := strings.Split(string(out), "\n")
	today := time.Now().Format("Jan 2")
	count := 0

	for _, line := range lines {
		if strings.Contains(line, today) {
			count++
		}
	}
	SSHFailedLoginsToday.Set(float64(count))
}
