package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"go.opentelemetry.io/otel/trace"

	infra "github.com/lcnascimento/istio-knative-poc/go-libs/infra"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/errors"

	"github.com/lcnascimento/istio-knative-poc/exports-service/domain"
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
	if in.Tracer == nil {
		return nil, fmt.Errorf("Missing required dependency: Tracer")
	}

	if in.Logger == nil {
		return nil, fmt.Errorf("Missing required dependency: Logger")
	}

	return &Service{in: in}, nil
}

// GetExport ...
func (s Service) GetExport(ctx context.Context, id string) (*domain.Export, error) {
	ctx, span := s.in.Tracer.Start(ctx, "domain.repository.GetExport")
	defer span.End()

	s.in.Logger.Info(ctx, "Fetching export %s from database", id)

	exports, err := s.ListExports(ctx)
	if err != nil {
		return nil, err
	}

	for _, export := range exports {
		if export.ID == id {
			return export, nil
		}
	}

	s.in.Logger.Error(ctx, errors.New(fmt.Sprintf("Could not find export %s in database", id)))
	return nil, domain.ErrEntityNotFound
}

// ListExports ...
func (s Service) ListExports(ctx context.Context) ([]*domain.Export, error) {
	ctx, span := s.in.Tracer.Start(ctx, "domain.repository.ListExports")
	defer span.End()

	s.in.Logger.Info(ctx, "Fetching exports from database")

	file, err := os.Open("config/exports.json")
	if err != nil {
		s.in.Logger.Error(ctx, errors.New(fmt.Sprintf("Could not list exports from database: %s", err.Error())))
		return nil, err
	}

	exports := []*domain.Export{}

	parser := json.NewDecoder(file)
	if err := parser.Decode(&exports); err != nil {
		s.in.Logger.Error(ctx, errors.New(fmt.Sprintf("Could not list exports from database: %s", err.Error())))
		return nil, err
	}

	return exports, nil
}
