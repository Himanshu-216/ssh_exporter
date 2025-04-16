package metrics

import (
	"os/exec"
	"strings"
	"time"
)

func UpdateLastLoginTimes() {
	out, err := exec.Command("lastlog").Output()
	if err != nil {
		return
	}
	lines := strings.Split(string(out), "\n")
	for _, line := range lines[1:] {
		fields := strings.Fields(line)
		if len(fields) < 8 {
			continue
		}
		user := fields[0]
		dateStr := strings.Join(fields[4:], " ")
		t, err := time.Parse("Mon Jan 2 15:04:05 -0700 2006", dateStr)
		if err == nil {
			LastLoginTime.WithLabelValues(user).Set(float64(t.Unix()))
		}
	}
}
