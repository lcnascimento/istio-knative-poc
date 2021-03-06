module github.com/lcnascimento/istio-knative-poc/audiences-service

go 1.14

replace github.com/lcnascimento/istio-knative-poc/exports-service => ../exports-service

replace github.com/lcnascimento/istio-knative-poc/segments-service => ../segments-service

replace github.com/lcnascimento/istio-knative-poc/go-libs => ../go-libs

require (
	github.com/golang/protobuf v1.4.3
	github.com/lcnascimento/istio-knative-poc/exports-service v0.0.0-00010101000000-000000000000
	github.com/lcnascimento/istio-knative-poc/go-libs v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/otel v0.16.0
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)
