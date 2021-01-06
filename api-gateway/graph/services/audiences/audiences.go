package audiences

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	pb "github.com/lcnascimento/istio-knative-poc/audiences-service/application/grpc/proto"

	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/model"
)

// ErrClientNotConnected ...
var ErrClientNotConnected = fmt.Errorf("client not connected to AudiencesService gRPC server")

// ServiceInput ...
type ServiceInput struct {
	ServerAddress string
}

// Service ...
type Service struct {
	in ServiceInput

	cli pb.AudiencesServiceFrontendClient
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

	s.cli = pb.NewAudiencesServiceFrontendClient(conn)

	return nil
}

// ListAudiences ...
func (s Service) ListAudiences(ctx context.Context) ([]*model.Audience, error) {
	if s.cli == nil {
		return nil, ErrClientNotConnected
	}

	res, err := s.cli.ListAudiences(ctx, &pb.ListAudiencesRequest{})
	if err != nil {
		return nil, err
	}

	audiences := []*model.Audience{}
	for _, audience := range res.Audiences {
		audiences = append(audiences, translate(audience))
	}

	return audiences, nil
}

// GetAudience ...
func (s Service) GetAudience(ctx context.Context, id string) (*model.Audience, error) {
	if s.cli == nil {
		return nil, ErrClientNotConnected
	}

	res, err := s.cli.GetAudience(ctx, &pb.GetAudienceRequest{
		AudienceId: id,
	})
	if err != nil {
		return nil, err
	}

	return translate(res.Audience), nil
}

var dtoToModelProvider = map[pb.AudienceProvider]model.AudienceProvider{
	pb.AudienceProvider_GOOGLE:   model.AudienceProviderGoogle,
	pb.AudienceProvider_FACEBOOK: model.AudienceProviderFacebook,
}

func translate(dto *pb.Audience) *model.Audience {
	return &model.Audience{
		ID:       dto.Id,
		AppKey:   dto.AppKey,
		Name:     dto.Name,
		Provider: dtoToModelProvider[dto.Provider],
		Segment: &model.Segment{
			ID: dto.SegmentId,
		},
		LastExport: &model.Export{
			ID: dto.LastExportId,
		},
	}
}
