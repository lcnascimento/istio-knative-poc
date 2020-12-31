package grpc

import (
	"context"
	"fmt"
	"log"

	pb "github.com/lcnascimento/istio-knative-poc/exports-service/application/grpc/proto"

	"github.com/lcnascimento/istio-knative-poc/exports-service/domain"
)

// WorkerInput ...
type WorkerInput struct {
	Exportator domain.Exportator
}

// Worker ...
type Worker struct {
	in WorkerInput

	pb.UnimplementedExportsServiceWorkerServer
}

// NewWorker ...
func NewWorker(in WorkerInput) (*Worker, error) {
	if in.Exportator == nil {
		return nil, fmt.Errorf("Missing Exportator dependency")
	}

	return &Worker{in: in}, nil
}

// ProcessExport ...
func (w Worker) ProcessExport(ctx context.Context, in *pb.ProcessExportRequest) (*pb.Void, error) {
	if err := w.in.Exportator.Export(ctx, in.ExportId); err != nil {
		log.Printf("Could not process exportation %s: %s", in.ExportId, err.Error())
		return nil, err
	}

	return &pb.Void{}, nil
}
