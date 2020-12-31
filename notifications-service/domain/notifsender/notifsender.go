package notifsender

import (
	"context"
	"fmt"
	"log"

	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain"
)

// SenderInput ...
type SenderInput struct {
	Repo      domain.NotificationsRepository
	Segments  domain.SegmentsService
	Providers map[domain.Channel]domain.NotificationProvider
}

// Sender ...
type Sender struct {
	in SenderInput
}

// NewSender ...
func NewSender(in SenderInput) (*Sender, error) {
	if in.Repo == nil {
		return nil, fmt.Errorf("Missing NotificationsRepository dependency")
	}

	if in.Segments == nil {
		return nil, fmt.Errorf("Missing SegmentsService dependency")
	}

	if in.Providers == nil {
		return nil, fmt.Errorf("Missing Providers dependency")
	}

	return &Sender{in: in}, nil
}

// SendNotification ...
func (s Sender) SendNotification(ctx context.Context, id string) error {
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
