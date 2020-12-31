package sender

import (
	"context"
	"fmt"
	"log"

	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain"
)

// ServiceInput ...
type ServiceInput struct {
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
	if in.Repo == nil {
		return nil, fmt.Errorf("Missing NotificationsRepository dependency")
	}

	if in.Segments == nil {
		return nil, fmt.Errorf("Missing SegmentsService dependency")
	}

	if in.Providers == nil {
		return nil, fmt.Errorf("Missing Providers dependency")
	}

	return &Service{in: in}, nil
}

// SendNotification ...
func (s Service) SendNotification(ctx context.Context, id string) error {
	log.Printf("Sending notification %s", id)

	notif, err := s.in.Repo.GetNotification(ctx, id)
	if err != nil {
		return err
	}

	provider := s.in.Providers[notif.Channel]
	if provider == nil {
		return fmt.Errorf("Unknown provider %s", notif.Channel)
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
