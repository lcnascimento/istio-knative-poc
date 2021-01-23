package sender

import (
	"context"
	"fmt"

	infra "github.com/lcnascimento/istio-knative-poc/go-libs/infra"
	"go.opentelemetry.io/otel/trace"

	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain"
)

// ServiceInput ...
type ServiceInput struct {
	Tracer    trace.Tracer
	Logger    infra.LogProvider
	Repo      domain.NotificationsRepository
	Segments  domain.SegmentsService
	Providers map[domain.Channel]domain.NotificationProvider
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

	if in.Repo == nil {
		return nil, fmt.Errorf("Missing required dependency: Repo")
	}

	if in.Segments == nil {
		return nil, fmt.Errorf("Missing dependency: Segments")
	}

	if in.Providers == nil {
		return nil, fmt.Errorf("Missing dependency: Providers")
	}

	return &Service{in: in}, nil
}

// SendNotification ...
func (s Service) SendNotification(ctx context.Context, id string) error {
	ctx, span := s.in.Tracer.Start(ctx, "domain.sender.SendNotification")
	defer span.End()

	s.in.Logger.Info(ctx, "Sending notification %s", id)

	notif, err := s.in.Repo.GetNotification(ctx, id)
	if err != nil {
		return err
	}

	provider := s.in.Providers[notif.Channel]
	if provider == nil {
		err := fmt.Errorf("Unknown provider %s", notif.Channel)
		s.in.Logger.Error(ctx, err)
		return err
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ch, errCh := s.in.Segments.GetSegmentUsers(ctx, notif.SegmentID)
	done, perrCh := provider.SendNotification(ctx, *notif, ch)

	select {
	case err := <-errCh:
		return err
	case err := <-perrCh:
		return err
	case <-done:
		return nil
	}
}
