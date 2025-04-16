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
			time.Sleep(15 * time.Second)
		}
	}()

	fmt.Println("Starting SSH Exporter on :2112")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
