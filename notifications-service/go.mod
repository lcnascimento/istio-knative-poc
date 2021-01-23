module github.com/lcnascimento/istio-knative-poc/notifications-service

replace github.com/lcnascimento/istio-knative-poc/go-libs => ../go-libs

replace github.com/lcnascimento/istio-knative-poc/segments-service => ../segments-service

go 1.14

require (
	github.com/golang/protobuf v1.4.3
	github.com/lcnascimento/istio-knative-poc/go-libs v0.0.0-00010101000000-000000000000
	github.com/lcnascimento/istio-knative-poc/segments-service v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc v0.11.0
	go.opentelemetry.io/otel v0.16.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.16.0
	go.opentelemetry.io/otel/sdk v0.16.0
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)
