package segments

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	pb "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc/proto"

	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/model"
)

// ServiceInput ...
type ServiceInput struct {
	ServerAddress string
}

// Service ...
type Service struct {
	in ServiceInput

	cli pb.SegmentsServiceFrontendClient
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

	s.cli = pb.NewSegmentsServiceFrontendClient(conn)

	return nil
}

// ListSegments ...
func (s Service) ListSegments(ctx context.Context) ([]*model.Segment, error) {
	if s.cli == nil {
		return nil, fmt.Errorf("client not connected to SegmentsService gRPC server")
	}

	res, err := s.cli.ListSegments(ctx, &pb.ListSegmentsRequest{})
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
	if s.cli == nil {
		return nil, fmt.Errorf("client not connected to SegmentsService gRPC server")
	}

	res, err := s.cli.GetSegment(ctx, &pb.GetSegmentRequest{
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
