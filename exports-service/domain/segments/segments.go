package segments

import (
	"context"
	"fmt"
	"io"

	"go.opentelemetry.io/otel/trace"

	infra "github.com/lcnascimento/istio-knative-poc/go-libs/infra"

	"github.com/lcnascimento/istio-knative-poc/exports-service/domain"
	pb "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc/proto"
)

// ServiceInput ...
type ServiceInput struct {
	BulkSize int

	Tracer trace.Tracer
	Logger infra.LogProvider
	Client pb.SegmentsServiceFrontendClient
}

// Service ...
type Service struct {
	in ServiceInput

	cli pb.SegmentsServiceFrontendClient
}

// NewService ...
func NewService(in ServiceInput) (*Service, error) {
	if in.Tracer == nil {
		return nil, fmt.Errorf("Missing required dependency: Tracer")
	}

	if in.Logger == nil {
		return nil, fmt.Errorf("Missing required dependency: Logger")
	}

	if in.BulkSize == 0 {
		in.BulkSize = 100
	}

	return &Service{in: in}, nil
}

// GetSegmentUsers ...
func (s Service) GetSegmentUsers(ctx context.Context, id string) (chan []*domain.User, chan error) {
	ctx, span := s.in.Tracer.Start(ctx, "domain.segments.GetSegmentUsers")
	defer span.End()

	s.in.Logger.Info(ctx, "Loading users from segment %s", id)

	ch := make(chan []*domain.User)
	errCh := make(chan error)

	go func() {
		defer close(ch)
		defer close(errCh)

		stream, err := s.in.Client.GetSegmentUsers(ctx, &pb.GetSegmentUsersRequest{
			SegmentId: id,
			Size:      int32(s.in.BulkSize),
		})
		if err != nil {
			errCh <- err
			return
		}

		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				s.in.Logger.Info(ctx, "All users from segment %s loaded successfuly", id)
				return
			}

			if err != nil {
				errCh <- err
				return
			}

			bulk := []*domain.User{}

			for _, dto := range resp.Data {
				user := domain.User{}
				user.FillByGRPCDTO(dto)

				bulk = append(bulk, &user)
				s.in.Logger.Info(ctx, "%s users loaded from segment %s", len(bulk), id)
			}

			ch <- bulk
		}
	}()

	return ch, errCh
}
