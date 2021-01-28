package metrics

import (
	"fmt"
	"net/http"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	prom "github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/contrib/instrumentation/host"
	"go.opentelemetry.io/otel/metric"

	prometheus "go.opentelemetry.io/otel/exporters/metric/prometheus"
)

// MeterInput ...
type MeterInput struct {
	ApplicationName  string
	ServerPort       int
	HistogramBuckets []float64
}

// Init ...
func Init(in MeterInput) (metric.MeterProvider, error) {
	expo, err := prometheus.InstallNewPipeline(prometheus.Config{})
	if err != nil {
		return nil, err
	}

	if err := host.Start(host.WithMeterProvider(expo.MeterProvider())); err != nil {
		return nil, err
	}

	labels := prom.Labels{
		"application": in.ApplicationName,
	}

	grpc_prometheus.WithConstLabels(labels)
	grpc_prometheus.EnableHandlingTimeHistogram(
		grpc_prometheus.WithHistogramBuckets(in.HistogramBuckets),
		grpc_prometheus.WithHistogramConstLabels(labels),
	)

	go func() {
		http.Handle("/metrics", expo)
		fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", in.ServerPort), nil))
	}()

	// TODO: Discover how to expose both handlers at the same path
	// go func() {
	// 	http.Handle("/metrics", promhttp.Handler())
	// 	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", in.ServerPort+1), nil))
	// }()

	return expo.MeterProvider(), nil
}
