package grpc

import (
	"context"
	"fmt"
	"log"

	pb "github.com/lcnascimento/istio-knative-poc/exports-service/application/grpc/proto"
	"github.com/lcnascimento/istio-knative-poc/exports-service/domain"
)

// FrontendInput ...
type FrontendInput struct {
	Repo     domain.ExportsRepository
	Enqueuer domain.ExportEnqueuer
}

// Frontend ...
type Frontend struct {
	in FrontendInput

	pb.UnimplementedExportsServiceFrontendServer
}

// NewFrontend ...
func NewFrontend(in FrontendInput) (*Frontend, error) {
	if in.Repo == nil {
		return nil, fmt.Errorf("Missing ExportsRepository dependency")
	}

	if in.Enqueuer == nil {
		return nil, fmt.Errorf("Missing ExportEnqueuer dependency")
	}

	return &Frontend{in: in}, nil
}

// GetExport ...
func (f Frontend) GetExport(ctx context.Context, in *pb.GetExportRequest) (*pb.GetExportResponse, error) {
	expo, err := f.in.Repo.GetExport(ctx, in.ExportId)
	if err != nil {
		log.Printf("Could not get export %s: %s", in.ExportId, err.Error())
		return nil, err
	}

	return &pb.GetExportResponse{Export: expo.ToGRPCDTO()}, nil
}

// ListExports ...
func (f Frontend) ListExports(ctx context.Context, _ *pb.ListExportsRequest) (*pb.ListExportsResponse, error) {
	expos, err := f.in.Repo.ListExports(ctx)
	if err != nil {
		log.Printf("Could not get exports: %s", err.Error())
		return nil, err
	}

	resp := &pb.ListExportsResponse{}
	for _, expo := range expos {
		resp.Exports = append(resp.Exports, expo.ToGRPCDTO())
	}

	return resp, nil
}

// EnqueueExport ...
func (f Frontend) EnqueueExport(ctx context.Context, in *pb.EnqueueExportRequest) (*pb.Void, error) {
	if err := f.in.Enqueuer.EnqueueExport(ctx, in.ExportId); err != nil {
		log.Printf("Could not enqueue export  %s: %s", in.ExportId, err.Error())
		return nil, err
	}

	return &pb.Void{}, nil
}
