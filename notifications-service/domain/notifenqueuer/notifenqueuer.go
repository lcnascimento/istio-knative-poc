package notifenqueuer

import (
	"context"

	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain"
)

// EnqueuerInput ...
type EnqueuerInput struct {
}

// Enqueuer ...
type Enqueuer struct {
}

// NewEnqueuer ...
func NewEnqueuer(in EnqueuerInput) (*Enqueuer, error) {
	return &Enqueuer{}, nil
}

// EnqueueNotification ...
func (e Enqueuer) EnqueueNotification(ctx context.Context, id string) error {
	return domain.ErrNotImplemented
}
