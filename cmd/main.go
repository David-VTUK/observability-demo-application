package main

import (
	"net/http"
	"time"

	"github.com/david-vtuk/observability-demo-application/internal/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {

	prometheus.Init()

	go func() {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			prometheus.ChangeValues()
			log.Info("Metrics updated")
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)

}
