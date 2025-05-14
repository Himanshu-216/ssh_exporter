package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	// Metric for tracking the number of active SSH sessions
	SSHConnections = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "ssh_active_sessions",
		Help: "Number of active SSH sessions",
	})

	// Metric for tracking the number of active SSH sessions per user
	SSHConnectionsByUser = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ssh_active_sessions_by_user",
			Help: "Number of active SSH sessions by user",
		},
		[]string{"user"},
	)

	// Metric for tracking the number of SSH logins today
	SSHLoginsToday = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "ssh_logins_today",
		Help: "Number of SSH logins today",
	})

	// Metric for tracking the last login time per user
	LastLoginTime = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ssh_user_last_login",
			Help: "Last login time per user (Unix timestamp)",
		},
		[]string{"user"},
	)

	// Metric for tracking failed SSH login attempts by IP address
	SSHFailedLoginsByIP = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ssh_failed_logins_by_ip",
			Help: "Number of failed SSH login attempts by IP address",
		},
		[]string{"ip"},
	)

	// Metric for tracking successful SSH login attempts by IP address
	SSHSuccessfulLoginsByIP = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ssh_successful_logins_by_ip",
			Help: "Number of successful SSH login attempts by IP address",
		},
		[]string{"ip"},
	)

	// Metric for tracking the total duration of SSH sessions
	SSHSessionDuration = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ssh_session_duration_seconds",
			Help: "Total time (in seconds) a user has been connected via SSH, from session start until removal",
		},
		[]string{"session_id", "user", "rsa_fingerprint"},
	)

	// Sessions map to store active sessions
	sessions = make(map[string]*Session)
)

// Register all metrics with Prometheus
func RegisterMetrics() {
	prometheus.MustRegister(SSHConnections)
	prometheus.MustRegister(SSHConnectionsByUser)
	prometheus.MustRegister(SSHLoginsToday)
	prometheus.MustRegister(LastLoginTime)
	prometheus.MustRegister(SSHFailedLoginsByIP)
	prometheus.MustRegister(SSHSuccessfulLoginsByIP)
	prometheus.MustRegister(SSHSessionDuration)
}
