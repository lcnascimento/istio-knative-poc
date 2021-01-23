package tracing

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// TracerInput ...
type TracerInput struct {
	ServiceName   string
	TracerName    string
	AgentEndpoint string
	Tags          []label.KeyValue
}

// Init ...
func Init(in TracerInput) (trace.Tracer, func(), error) {
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithAgentEndpoint(in.AgentEndpoint),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: in.ServiceName,
			Tags:        in.Tags,
		}),
		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
	)
	if err != nil {
		return nil, nil, err
	}

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return otel.Tracer(in.TracerName), flush, nil
}
