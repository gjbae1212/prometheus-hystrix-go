package main

import (
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	hystrix_metric "github.com/afex/hystrix-go/hystrix/metric_collector"
	prometheus_hystrix_go "github.com/gjbae1212/prometheus-hystrix-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// get hystrix wrapper
	wrapper := prometheus_hystrix_go.NewPrometheusCollector("hystrix", map[string]string{"app": "test"})

	// register and initialize to hystrix prometheus
	hystrix_metric.Registry.Register(wrapper)

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		hystrix.Do("hystrix-test", func() error { return nil }, nil)
	})

	// start server
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
