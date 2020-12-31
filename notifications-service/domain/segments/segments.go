package segments

import (
	"context"
	"fmt"
	"io"

	"google.golang.org/grpc"

	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain"
	pb "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc/proto"
)

// ServiceInput ...
type ServiceInput struct {
	ServerAddress string
}

// Service ...
type Service struct {
	in ServiceInput

	cli pb.SegmentsServiceClient
}

// NewService ...
func NewService(in ServiceInput) (*Service, error) {
	return &Service{in: in}, nil
}

// Connect ...
func (s *Service) Connect() error {
	conn, err := grpc.Dial(s.in.ServerAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}

	s.cli = pb.NewSegmentsServiceClient(conn)

	return nil
}

// GetSegmentUsers ...
func (s Service) GetSegmentUsers(ctx context.Context, id string) (chan []*domain.User, chan error) {
	ch := make(chan []*domain.User)
	errCh := make(chan error)

	go func() {
		defer close(ch)
		defer close(errCh)

		if s.cli == nil {
			errCh <- fmt.Errorf("client not connected to gRPC server")
			return
		}

		stream, err := s.cli.GetSegmentUsers(ctx, &pb.GetSegmentUsersRequest{SegmentId: id, Size: 1})
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
