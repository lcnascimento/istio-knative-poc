package enqueuer

import (
	"context"

	"github.com/lcnascimento/istio-knative-poc/audiences-service/domain"
)

// ServiceInput ...
type ServiceInput struct{}

// Service ...
type Service struct {
	in ServiceInput
}

// NewService ...
func NewService(in ServiceInput) (*Service, error) {
	return &Service{in: in}, nil
}

// EnqueueAudienceSending ...
func (s Service) EnqueueAudienceSending(ctx context.Context, id string) error {
	return domain.ErrNotImplemented
}
