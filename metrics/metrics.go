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

	SSHConnectionsByFingerprint = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ssh_connections_by_fingerprint",
			Help: "SSH connection count by fingerprint",
		},
		[]string{"fingerprint"},
	)

	SSHFailedLoginsByIP = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ssh_failed_logins_by_ip",
			Help: "Number of failed SSH login attempts by IP address",
		},
		[]string{"ip"},
	)

	SSHSuccessfulLoginsByIP = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ssh_successful_logins_by_ip",
			Help: "Number of successful SSH login attempts by IP address",
		},
		[]string{"ip"},
	)
)

func RegisterMetrics() {
	prometheus.MustRegister(SSHConnections)
	prometheus.MustRegister(SSHConnectionsByUser)
	prometheus.MustRegister(SSHLoginsToday)
	prometheus.MustRegister(LastLoginTime)
	prometheus.MustRegister(SSHConnectionsByFingerprint)
	prometheus.MustRegister(SSHFailedLoginsByIP)
	prometheus.MustRegister(SSHSuccessfulLoginsByIP)
}
