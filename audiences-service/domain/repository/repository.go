package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"go.opentelemetry.io/otel/trace"

	infra "github.com/lcnascimento/istio-knative-poc/go-libs/infra"

	"github.com/lcnascimento/istio-knative-poc/audiences-service/domain"
)

// ServiceInput ...
type ServiceInput struct {
	Tracer trace.Tracer
	Logger infra.LogProvider
}

// Service ...
type Service struct {
	in ServiceInput
}

// NewService ...
func NewService(in ServiceInput) (*Service, error) {
	if in.Logger == nil {
		return nil, fmt.Errorf("Missing required dependency: Logger")
	}

	if in.Tracer == nil {
		return nil, fmt.Errorf("Missing required dependency: Tracer")
	}

	return &Service{in: in}, nil
}

// GetAudience ...
func (s Service) GetAudience(ctx context.Context, id string) (*domain.Audience, error) {
	ctx, span := s.in.Tracer.Start(ctx, "domain.repository.GetAudience")
	defer span.End()

	s.in.Logger.Info(ctx, "Fetching audience %s from database", id)

	audiences, err := s.ListAudiences(ctx)
	if err != nil {
		return nil, err
	}

	for _, audience := range audiences {
		if audience.ID == id {
			return audience, nil
		}
	}

	s.in.Logger.Error(ctx, fmt.Errorf("Could not find audience %s", id))
	return nil, domain.ErrEntityNotFound
}

// ListAudiences ...
func (s Service) ListAudiences(ctx context.Context) ([]*domain.Audience, error) {
	ctx, span := s.in.Tracer.Start(ctx, "domain.repository.ListAudiences")
	defer span.End()

	s.in.Logger.Info(ctx, "Fetching audiences from database")

	file, err := os.Open("config/audiences.json")
	if err != nil {
		s.in.Logger.Error(ctx, fmt.Errorf("Could not load audiences from database: %s", err.Error()))
		return nil, err
	}

	audiences := []*domain.Audience{}

	parser := json.NewDecoder(file)
	if err := parser.Decode(&audiences); err != nil {
		s.in.Logger.Error(ctx, fmt.Errorf("Could not load audiences from database: %s", err.Error()))
		return nil, err
	}

	return audiences, nil
}
