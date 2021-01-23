package exports

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"

	pb "github.com/lcnascimento/istio-knative-poc/exports-service/application/grpc/proto"

	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/model"
	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/services"
)

// ServiceInput ...
type ServiceInput struct {
	Tracer trace.Tracer
	Client pb.ExportsServiceFrontendClient
}

// Service ...
type Service struct {
	in ServiceInput
}

// NewService ...
func NewService(in ServiceInput) (*Service, error) {
	if in.Tracer == nil {
		return nil, fmt.Errorf("Missing required dependency: Tracer")
	}

	if in.Client == nil {
		return nil, fmt.Errorf("Missing required dependency: Client")
	}

	return &Service{in: in}, nil
}

// ListExports ...
func (s Service) ListExports(ctx context.Context) ([]*model.Export, error) {
	ctx, span := s.in.Tracer.Start(ctx, "graph.services.exports.ListExports")
	defer span.End()

	res, err := s.in.Client.ListExports(ctx, &pb.ListExportsRequest{})
	if err != nil {
		return nil, err
	}

	exports := []*model.Export{}
	for _, export := range res.Exports {
		exports = append(exports, translate(export))
	}

	return exports, nil
}

// GetExport ...
func (s Service) GetExport(ctx context.Context, id string) (*model.Export, error) {
	ctx, span := s.in.Tracer.Start(ctx, "graph.services.exports.GetExport")
	defer span.End()

	res, err := s.in.Client.GetExport(ctx, &pb.GetExportRequest{
		ExportId: id,
	})
	if err != nil {
		return nil, err
	}

	return translate(res.Export), nil
}

// CreateExport ...
func (s Service) CreateExport(ctx context.Context, in model.NewExport) (*model.Export, error) {
	ctx, span := s.in.Tracer.Start(ctx, "graph.services.exports.CreateExport")
	defer span.End()

	return nil, services.ErrNotImplemented
}

var dtoToModelModule = map[pb.ExportModule]model.ExportModule{
	pb.ExportModule_USERS: model.ExportModuleUsers,
	pb.ExportModule_ADS:   model.ExportModuleAds,
}

var dtoToModelStatus = map[pb.ExportStatus]model.JobStatus{
	pb.ExportStatus_CREATED: model.JobStatusCreated,
	pb.ExportStatus_RUNNING: model.JobStatusRunning,
	pb.ExportStatus_FAILED:  model.JobStatusFailed,
	pb.ExportStatus_DONE:    model.JobStatusDone,
}

func translate(dto *pb.Export) *model.Export {
	return &model.Export{
		ID:     dto.Id,
		AppKey: dto.AppKey,
		Name:   dto.Name,
		Module: dtoToModelModule[dto.Module],
		Status: dtoToModelStatus[dto.Status],
		Segment: &model.Segment{
			ID: dto.SegmentId,
		},
	}
}
