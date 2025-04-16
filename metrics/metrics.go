package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	SSHConnections = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "ssh_active_sessions",
		Help: "Number of active SSH sessions",
	})

	SSHConnectionsByUser = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ssh_active_sessions_by_user",
			Help: "Number of active SSH sessions by user",
		},
		[]string{"user"},
	)

	SSHLoginsToday = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "ssh_logins_today",
		Help: "Number of SSH logins today",
	})

	LastLoginTime = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ssh_user_last_login",
			Help: "Last login time per user (Unix timestamp)",
		},
		[]string{"user"},
	)
)

func RegisterMetrics() {
	prometheus.MustRegister(SSHConnections)
	prometheus.MustRegister(SSHConnectionsByUser)
	prometheus.MustRegister(SSHLoginsToday)
	prometheus.MustRegister(LastLoginTime)
}
