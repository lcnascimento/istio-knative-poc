package grpc

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	"go.opentelemetry.io/otel/trace"

	infra "github.com/lcnascimento/istio-knative-poc/go-libs/infra"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/errors"

	pb "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc/proto"
	"github.com/lcnascimento/istio-knative-poc/segments-service/domain"
)

// FrontendInput ...
type FrontendInput struct {
	Tracer trace.Tracer
	Repo   domain.SegmentsRepository
	Logger infra.LogProvider
}

// Frontend ...
type Frontend struct {
	in FrontendInput

	pb.UnimplementedSegmentsServiceFrontendServer
}

// NewFrontend ...
func NewFrontend(in FrontendInput) (*Frontend, error) {
	if in.Tracer == nil {
		return nil, fmt.Errorf("Missing required dependency: Tracer")
	}

	if in.Logger == nil {
		return nil, fmt.Errorf("Missing required dependency: Logger")
	}

	if in.Repo == nil {
		return nil, fmt.Errorf("Missing required dependency: Repo")
	}

	return &Frontend{in: in}, nil
}

// GetSegment ...
func (s Frontend) GetSegment(ctx context.Context, in *pb.GetSegmentRequest) (*pb.GetSegmentResponse, error) {
	ctx, span := s.in.Tracer.Start(ctx, "application.grpc.GetSegment")
	defer span.End()

	segment, err := s.in.Repo.FindSegment(ctx, in.SegmentId)
	if err != nil {
		return nil, err
	}

	return &pb.GetSegmentResponse{Segment: segment.ToGRPCDTO()}, nil
}

// ListSegments ...
func (s Frontend) ListSegments(ctx context.Context, in *pb.ListSegmentsRequest) (*pb.ListSegmentsResponse, error) {
	ctx, span := s.in.Tracer.Start(ctx, "application.grpc.ListSegments")
	defer span.End()

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
	ctx, span := s.in.Tracer.Start(srv.Context(), "application.grpc.GetSegmentUsers")
	defer span.End()

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
				s.in.Logger.Error(ctx, errors.New(fmt.Sprintf("Something went wrong with in the connection context: %s", err.Error())))
				done()
				return err
			}

			if err := srv.Send(&resp); err != nil {
				s.in.Logger.Error(ctx, errors.New(fmt.Sprintf("Could not send package to gRPC client: %s", err.Error())))
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
		return err
	}

	return nil
}
