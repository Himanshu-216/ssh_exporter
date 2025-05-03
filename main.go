package main

import (
	"fmt"
	"github.com/Himanshu-216/ssh_exporter/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func main() {
	metrics.RegisterMetrics()

	go func() {
		for {
			metrics.UpdateSSHConnections()
			metrics.UpdateLoginsToday()
			metrics.UpdateLastLoginTimes()
			metrics.LoginsmonitorwithIP()
			time.Sleep(5 * time.Second)
		}
	}()

	fmt.Println("Starting SSH Exporter on :9898")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9898", nil)
}
