package main

import (
	"net/http"

	"https://github.com/FullStackS-GmbH/trx-metrics-exporter/pkg/collector"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	// This section will start the HTTP Server and expose
	// any metrics on the /metrics endpoint.
	http.Handle("/metrics", promhttp.Handler())
	log.Info("Beginning to serve on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
