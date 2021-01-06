module github.com/lcnascimento/istio-knative-poc/api-gateway

replace github.com/lcnascimento/istio-knative-poc/go-libs => ../go-libs

replace github.com/lcnascimento/istio-knative-poc/segments-service => ../segments-service

replace github.com/lcnascimento/istio-knative-poc/notifications-service => ../notifications-service

replace github.com/lcnascimento/istio-knative-poc/exports-service => ../exports-service

replace github.com/lcnascimento/istio-knative-poc/audiences-service => ../audiences-service

go 1.14

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/lcnascimento/istio-knative-poc/audiences-service v0.0.0-00010101000000-000000000000
	github.com/lcnascimento/istio-knative-poc/exports-service v0.0.0-00010101000000-000000000000
	github.com/lcnascimento/istio-knative-poc/go-libs v0.0.0-00010101000000-000000000000
	github.com/lcnascimento/istio-knative-poc/notifications-service v0.0.0-00010101000000-000000000000
	github.com/lcnascimento/istio-knative-poc/segments-service v0.0.0-00010101000000-000000000000
	github.com/vektah/gqlparser/v2 v2.1.0
	google.golang.org/grpc v1.34.0
)
