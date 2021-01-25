package firebase

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"

	infra "github.com/lcnascimento/istio-knative-poc/go-libs/infra"

	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain"
)

// ServiceInput ...
type ServiceInput struct {
	Logger infra.LogProvider
	Tracer trace.Tracer
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

// SendNotification ...
func (s Service) SendNotification(ctx context.Context, notif domain.Notification, ch chan []*domain.User) (chan bool, chan error) {
	ctx, span := s.in.Tracer.Start(ctx, "domain.firebase.SendNotification")

	done := make(chan bool)
	errCh := make(chan error)

	go func() {
		defer span.End()

		for bulk := range ch {
			s.in.Logger.Info(ctx, "Sending %d WebPushs via Firebase", len(bulk))
		}

		done <- true
		close(done)
		close(errCh)
	}()

	return done, errCh
}
