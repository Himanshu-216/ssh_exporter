package main

import (
	"flag"
	"fmt"
	"github.com/Himanshu-216/ssh_exporter/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func main() {
	port := flag.String("web.listen-address", "9898", "Port on which to expose metrics and web interface (without ':')")
	flag.Parse()
	address := ":" + *port

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

	fmt.Printf("Starting SSH Exporter on %s\n", address)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(address, nil)
}
