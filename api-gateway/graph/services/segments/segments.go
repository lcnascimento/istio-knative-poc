package segments

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"

	pb "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc/proto"

	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/model"
)

// ServiceInput ...
type ServiceInput struct {
	Tracer trace.Tracer
	Client pb.SegmentsServiceFrontendClient
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

// ListSegments ...
func (s Service) ListSegments(ctx context.Context) ([]*model.Segment, error) {
	ctx, span := s.in.Tracer.Start(ctx, "graph.services.segments.ListSegments")
	defer span.End()

	res, err := s.in.Client.ListSegments(ctx, &pb.ListSegmentsRequest{})
	if err != nil {
		return nil, err
	}

	segments := []*model.Segment{}
	for _, seg := range res.Segments {
		segments = append(segments, translate(seg))
	}

	return segments, nil
}

// GetSegment ...
func (s Service) GetSegment(ctx context.Context, id string) (*model.Segment, error) {
	ctx, span := s.in.Tracer.Start(ctx, "graph.services.segments.GetSegment")
	defer span.End()

	res, err := s.in.Client.GetSegment(ctx, &pb.GetSegmentRequest{
		SegmentId: id,
	})
	if err != nil {
		return nil, err
	}

	return translate(res.Segment), nil
}

func translate(dto *pb.Segment) *model.Segment {
	var description *string
	if dto.Description != "" {
		description = &dto.Description
	}

	return &model.Segment{
		ID:          dto.Id,
		AppKey:      dto.AppKey,
		Name:        dto.Name,
		Description: description,
		Version:     int(dto.Version),
	}
}
