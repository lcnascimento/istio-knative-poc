package exports

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	pb "github.com/lcnascimento/istio-knative-poc/exports-service/application/grpc/proto"

	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/model"
	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/services"
)

// ServiceInput ...
type ServiceInput struct {
	ServerAddress string
}

// Service ...
type Service struct {
	in ServiceInput

	cli pb.ExportsServiceFrontendClient
}

// NewService ...
func NewService(in ServiceInput) (*Service, error) {
	if in.ServerAddress == "" {
		return nil, fmt.Errorf("Missing required dependency: ServerAddress")
	}

	return &Service{in: in}, nil
}

// Connect ...
func (s *Service) Connect() error {
	conn, err := grpc.Dial(s.in.ServerAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}

	s.cli = pb.NewExportsServiceFrontendClient(conn)

	return nil
}

// ListExports ...
func (s Service) ListExports(ctx context.Context) ([]*model.Export, error) {
	if s.cli == nil {
		return nil, fmt.Errorf("client not connected to ExportsService gRPC server")
	}

	res, err := s.cli.ListExports(ctx, &pb.ListExportsRequest{})
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
	if s.cli == nil {
		return nil, fmt.Errorf("client not connected to ExportsService gRPC server")
	}

	res, err := s.cli.GetExport(ctx, &pb.GetExportRequest{
		ExportId: id,
	})
	if err != nil {
		return nil, err
	}

	return translate(res.Export), nil
}

// CreateExport ...
func (s Service) CreateExport(ctx context.Context, in model.NewExport) (*model.Export, error) {
	if s.cli == nil {
		return nil, fmt.Errorf("client not connected to ExportsService gRPC server")
	}

	return nil, services.ErrNotImplemented
}

func translate(dto *pb.Export) *model.Export {
	return &model.Export{
		ID:     dto.Id,
		AppKey: dto.AppKey,
		Name:   dto.Name,
		// Module: dto.Module,
		// Segment: ,
		// Status: ,
	}
}
