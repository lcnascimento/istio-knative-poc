package grpc

import (
	"context"
	"log"

	pb "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc/proto"
	"github.com/lcnascimento/istio-knative-poc/segments-service/domain"
	"golang.org/x/sync/errgroup"
)

// FrontendInput ...
type FrontendInput struct {
	Repo domain.SegmentsRepository
}

// Frontend ...
type Frontend struct {
	in FrontendInput

	pb.UnimplementedSegmentsServiceFrontendServer
}

// NewFrontend ...
func NewFrontend(in FrontendInput) *Frontend {
	return &Frontend{in: in}
}

// GetSegment ...
func (s Frontend) GetSegment(ctx context.Context, in *pb.GetSegmentRequest) (*pb.GetSegmentResponse, error) {
	segment, err := s.in.Repo.FindSegment(ctx, in.SegmentId)
	if err != nil {
		return nil, err
	}

	return &pb.GetSegmentResponse{Segment: segment.ToGRPCDTO()}, nil
}

// ListSegments ...
func (s Frontend) ListSegments(ctx context.Context, in *pb.ListSegmentsRequest) (*pb.ListSegmentsResponse, error) {
	segments, err := s.in.Repo.ListSegments(ctx)
	if err != nil {
		return nil, err
	}

	res := []*pb.Segment{}
	for _, seg := range segments {
		res = append(res, seg.ToGRPCDTO())
	}

	return &pb.ListSegmentsResponse{Segments: res}, nil
}

// GetSegmentUsers ...
func (s Frontend) GetSegmentUsers(in *pb.GetSegmentUsersRequest, srv pb.SegmentsServiceFrontend_GetSegmentUsersServer) error {
	ctx, done := context.WithCancel(srv.Context())
	defer done()

	ch, errCh := s.in.Repo.GetSegmentUsers(ctx, in.SegmentId, domain.GetSegmentUsersOptions{
		BulkSize: int(in.Size),
	})

	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		for users := range ch {
			resp := pb.GetSegmentUsersResponse{}
			for _, u := range users {
				resp.Data = append(resp.Data, u.ToGRPCDTO())
			}

			if err := srv.Context().Err(); err != nil {
				done()
				return err
			}

			log.Println("Sending message to client")
			if err := srv.Send(&resp); err != nil {
				done()
				return err
			}
		}

		return nil
	})

	g.Go(func() error {
		for err := range errCh {
			done()
			return err
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		log.Printf("Something went wrong: %s", err.Error())
		return err
	}

	log.Println("Segment's users transfer completed")
	return nil
}
