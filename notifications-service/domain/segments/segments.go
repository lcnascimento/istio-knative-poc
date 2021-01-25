package segments

import (
	"context"
	"fmt"
	"io"

	"go.opentelemetry.io/otel/trace"

	infra "github.com/lcnascimento/istio-knative-poc/go-libs/infra"

	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain"
	pb "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc/proto"
)

// ServiceInput ...
type ServiceInput struct {
	ServerAddress string
	BulkSize      int

	Tracer trace.Tracer
	Logger infra.LogProvider
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

	if in.Logger == nil {
		return nil, fmt.Errorf("Missing required dependency: Logger")
	}

	if in.Client == nil {
		return nil, fmt.Errorf("Missing required dependency: Client")
	}

	return &Service{in: in}, nil
}

// GetSegmentUsers ...
func (s Service) GetSegmentUsers(ctx context.Context, id string) (chan []*domain.User, chan error) {
	ctx, span := s.in.Tracer.Start(ctx, "domain.segments.GetSegmentUsers")

	s.in.Logger.Info(ctx, "Streaming users from segment %s", id)

	ch := make(chan []*domain.User)
	errCh := make(chan error)

	go func() {
		defer span.End()
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
				return
			}

			if err != nil {
				err := fmt.Errorf("failed trying to retrieve data from gRPC server stream")
				s.in.Logger.Error(ctx, err)
				errCh <- err
				return
			}

			bulk := []*domain.User{}

			for _, dto := range resp.Data {
				user := domain.User{}
				user.FillByGRPCDTO(dto)

				bulk = append(bulk, &user)
			}

			ch <- bulk
		}
	}()

	return ch, errCh
}
