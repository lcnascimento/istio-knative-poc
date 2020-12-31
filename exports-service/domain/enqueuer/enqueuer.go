package enqueuer

import (
	"context"

	"github.com/lcnascimento/istio-knative-poc/exports-service/domain"
)

// ServiceInput ...
type ServiceInput struct {
}

// Service ...
type Service struct {
	in ServiceInput
}

// NewService ...
func NewService(in ServiceInput) (*Service, error) {
	return &Service{in: in}, nil
}

// EnqueueExport ...
func (s Service) EnqueueExport(ctx context.Context, id string) error {
	return domain.ErrNotImplemented
}
