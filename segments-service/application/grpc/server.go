package grpc

import (
	"context"
	"log"
	"sync"

	pb "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc/proto"
	"github.com/lcnascimento/istio-knative-poc/segments-service/domain"
)

// ServerInput ...
type ServerInput struct {
	Repo domain.SegmentsRepository
}

// Server ...
type Server struct {
	in ServerInput

	pb.UnimplementedSegmentsServiceServer
}

// NewServer ...
func NewServer(in ServerInput) *Server {
	return &Server{in: in}
}

// GetSegment ...
func (s Server) GetSegment(ctx context.Context, in *pb.GetSegmentRequest) (*pb.GetSegmentResponse, error) {
	segment, err := s.in.Repo.FindSegment(ctx, in.SegmentId)
	if err != nil {
		return nil, err
	}

	return &pb.GetSegmentResponse{Segment: segment.ToGRPCDTO()}, nil
}

// GetSegmentUsers ...
func (s Server) GetSegmentUsers(in *pb.GetSegmentUsersRequest, srv pb.SegmentsService_GetSegmentUsersServer) error {
	log.Println("Fetching users from segment ", in.SegmentId)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		ch, err := s.in.Repo.GetSegmentUsers(srv.Context(), in.SegmentId, domain.GetSegmentUsersOptions{
			BulkSize: int(in.Size),
		})
		if err != nil {
			log.Fatalf("could not create GetSegmentUsers stream: %v", err)
		}

		for users := range ch {
			resp := pb.GetSegmentUsersResponse{}
			for _, u := range users {
				resp.Data = append(resp.Data, u.ToGRPCDTO())
			}

			log.Println("Sending message to client")
			if err := srv.Send(&resp); err != nil {
				log.Println("could not send message to client: ", err)
			}
		}

		log.Println("All messages sent to client")

		wg.Done()
	}()

	wg.Wait()

	return nil
}
