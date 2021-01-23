package audiences

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"

	pb "github.com/lcnascimento/istio-knative-poc/audiences-service/application/grpc/proto"

	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/model"
)

// ServiceInput ...
type ServiceInput struct {
	Tracer trace.Tracer
	Client pb.AudiencesServiceFrontendClient
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

	return &Service{in: in}, nil
}

// ListAudiences ...
func (s Service) ListAudiences(ctx context.Context) ([]*model.Audience, error) {
	ctx, span := s.in.Tracer.Start(ctx, "graph.services.audiences.ListAudiences")
	defer span.End()

	res, err := s.in.Client.ListAudiences(ctx, &pb.ListAudiencesRequest{})
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
	ctx, span := s.in.Tracer.Start(ctx, "graph.services.audiences.GetAudience")
	defer span.End()

	res, err := s.in.Client.GetAudience(ctx, &pb.GetAudienceRequest{
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
