package main

import (
	"net/http"

	hystrix_metric "github.com/afex/hystrix-go/hystrix/metric_collector"
	prometheus_hystrix_go "github.com/gjbae1212/prometheus-hystrix-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// get hystrix wrapper
	wrapper := prometheus_hystrix_go.NewPrometheusCollector("hystrix", "circuit", map[string]string{"app": "test"})

	// register and initialize to hystrix prometheus
	hystrix_metric.Registry.Register(wrapper)
	hystrix_metric.Registry.InitializeMetricCollectors("")

	// start server
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
