package enqueuer

import (
	"context"
	"fmt"

	infra "github.com/lcnascimento/istio-knative-poc/go-libs/infra"
	"go.opentelemetry.io/otel/trace"

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

// EnqueueExport ...
func (s Service) EnqueueExport(ctx context.Context, id string) error {
	s.in.Logger.Error(ctx, domain.ErrNotImplemented)
	return domain.ErrNotImplemented
}
