package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"ssh_exporter/metrics"
)

func main() {
	metrics.RegisterMetrics()

	go func() {
		for {
			metrics.UpdateSSHConnections()
			metrics.UpdateLoginsToday()
			metrics.UpdateLastLoginTimes()
			metrics.UpdateFailedLogins() 
			time.Sleep(5 * time.Second)
		}
	}()

	fmt.Println("Starting SSH Exporter on :9898")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9898", nil)
}
