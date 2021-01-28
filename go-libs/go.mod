module github.com/lcnascimento/istio-knative-poc/go-libs

go 1.14

require (
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/ditointernet/go-dito-errors v0.0.0-20200921194425-daad8d680a71
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.9.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc v0.11.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.16.0
	go.opentelemetry.io/contrib/instrumentation/host v0.16.0
	go.opentelemetry.io/otel v0.16.0
	go.opentelemetry.io/otel/exporters/metric/prometheus v0.16.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.16.0
	go.opentelemetry.io/otel/sdk v0.16.0
	go.uber.org/atomic v1.7.0 // indirect
	google.golang.org/grpc v1.34.0
)
