package grpc

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/lcnascimento/istio-knative-poc/audiences-service/domain"

	pb "github.com/lcnascimento/istio-knative-poc/audiences-service/application/grpc/proto"
)

// FrontendInput ...
type FrontendInput struct {
	Tracer   trace.Tracer
	Repo     domain.AudiencesRepository
	Enqueuer domain.AudienceSendingEnqueuer
}

// Frontend ...
type Frontend struct {
	in FrontendInput

	pb.UnimplementedAudiencesServiceFrontendServer
}

// NewFrontend ...
func NewFrontend(in FrontendInput) (*Frontend, error) {
	if in.Tracer == nil {
		return nil, fmt.Errorf("Missing required dependency: Tracer")
	}

	if in.Repo == nil {
		return nil, fmt.Errorf("Missing required dependency: Repo")
	}

	if in.Enqueuer == nil {
		return nil, fmt.Errorf("Missing required dependency: Enqueuer")
	}

	return &Frontend{in: in}, nil
}

// GetAudience ...
func (f Frontend) GetAudience(ctx context.Context, in *pb.GetAudienceRequest) (*pb.GetAudienceResponse, error) {
	ctx, span := f.in.Tracer.Start(ctx, "application.grpc.frontend.GetAudience")
	defer span.End()

	audience, err := f.in.Repo.GetAudience(ctx, in.AudienceId)
	if err != nil {
		return nil, err
	}

	return &pb.GetAudienceResponse{Audience: audience.ToGRPCDTO()}, nil
}

// ListAudiences ...
func (f Frontend) ListAudiences(ctx context.Context, _ *pb.ListAudiencesRequest) (*pb.ListAudiencesResponse, error) {
	ctx, span := f.in.Tracer.Start(ctx, "application.grpc.frontend.ListAudiences")
	defer span.End()

	audiences, err := f.in.Repo.ListAudiences(ctx)
	if err != nil {
		return nil, err
	}

	response := []*pb.Audience{}
	for _, aud := range audiences {
		response = append(response, aud.ToGRPCDTO())
	}

	return &pb.ListAudiencesResponse{Audiences: response}, nil
}

// EnqueueAudienceSending ...
func (f Frontend) EnqueueAudienceSending(ctx context.Context, in *pb.EnqueueAudienceSendingRequest) (*wrapperspb.BoolValue, error) {
	ctx, span := f.in.Tracer.Start(ctx, "application.grpc.frontend.EnqueueAudienceSending")
	defer span.End()

	if err := f.in.Enqueuer.EnqueueAudienceSending(ctx, in.AudienceId); err != nil {
		return wrapperspb.Bool(false), err
	}

	return wrapperspb.Bool(true), nil
}
