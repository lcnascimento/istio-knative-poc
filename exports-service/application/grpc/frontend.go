package grpc

import (
	"context"
	"fmt"

	pb "github.com/lcnascimento/istio-knative-poc/exports-service/application/grpc/proto"
	"go.opentelemetry.io/otel/trace"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/lcnascimento/istio-knative-poc/exports-service/domain"
)

// FrontendInput ...
type FrontendInput struct {
	Tracer   trace.Tracer
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
	if in.Tracer == nil {
		return nil, fmt.Errorf("Missing required dependency: Tracer")
	}

	if in.Repo == nil {
		return nil, fmt.Errorf("Missing required dependency: Repo")
	}

	if in.Enqueuer == nil {
		return nil, fmt.Errorf("Missing required dependency: Enqueuer")
	}

	return &Frontend{in: in}, nil
}

// GetExport ...
func (f Frontend) GetExport(ctx context.Context, in *pb.GetExportRequest) (*pb.GetExportResponse, error) {
	ctx, span := f.in.Tracer.Start(ctx, "application.grpc.GetExport")
	defer span.End()

	expo, err := f.in.Repo.GetExport(ctx, in.ExportId)
	if err != nil {
		return nil, err
	}

	return &pb.GetExportResponse{Export: expo.ToGRPCDTO()}, nil
}

// ListExports ...
func (f Frontend) ListExports(ctx context.Context, _ *pb.ListExportsRequest) (*pb.ListExportsResponse, error) {
	ctx, span := f.in.Tracer.Start(ctx, "application.grpc.ListExports")
	defer span.End()

	expos, err := f.in.Repo.ListExports(ctx)
	if err != nil {
		return nil, err
	}

	resp := &pb.ListExportsResponse{}
	for _, expo := range expos {
		resp.Exports = append(resp.Exports, expo.ToGRPCDTO())
	}

	return resp, nil
}

// EnqueueExport ...
func (f Frontend) EnqueueExport(ctx context.Context, in *pb.EnqueueExportRequest) (*wrapperspb.BoolValue, error) {
	ctx, span := f.in.Tracer.Start(ctx, "application.grpc.EnqueueExport")
	defer span.End()

	if err := f.in.Enqueuer.EnqueueExport(ctx, in.ExportId); err != nil {
		return wrapperspb.Bool(false), err
	}

	return wrapperspb.Bool(true), nil
}
