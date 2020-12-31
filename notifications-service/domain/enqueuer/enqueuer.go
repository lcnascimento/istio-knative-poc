package enqueuer

import (
	"context"

	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain"
)

// ServiceInput ...
type ServiceInput struct {
}

// Service ...
type Service struct {
}

// NewService ...
func NewService(in ServiceInput) (*Service, error) {
	return &Service{}, nil
}

// EnqueueNotification ...
func (e Service) EnqueueNotification(ctx context.Context, id string) error {
	return domain.ErrNotImplemented
}
