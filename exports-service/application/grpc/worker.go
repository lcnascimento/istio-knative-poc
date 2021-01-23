package grpc

import (
	"context"
	"fmt"

	pb "github.com/lcnascimento/istio-knative-poc/exports-service/application/grpc/proto"
	"go.opentelemetry.io/otel/trace"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/lcnascimento/istio-knative-poc/exports-service/domain"
)

// WorkerInput ...
type WorkerInput struct {
	Tracer     trace.Tracer
	Exportator domain.Exportator
}

// Worker ...
type Worker struct {
	in WorkerInput

	pb.UnimplementedExportsServiceWorkerServer
}

// NewWorker ...
func NewWorker(in WorkerInput) (*Worker, error) {
	if in.Tracer == nil {
		return nil, fmt.Errorf("Missing required dependency: Tracer")
	}

	if in.Exportator == nil {
		return nil, fmt.Errorf("Missing required dependency: Exportator")
	}

	return &Worker{in: in}, nil
}

// ProcessExport ...
func (w Worker) ProcessExport(ctx context.Context, in *pb.ProcessExportRequest) (*wrapperspb.BoolValue, error) {
	ctx, span := w.in.Tracer.Start(ctx, "application.grpc.ProcessExport")
	defer span.End()

	if err := w.in.Exportator.Export(ctx, in.ExportId); err != nil {
		return wrapperspb.Bool(false), err
	}

	return wrapperspb.Bool(true), nil
}
