package sendgrid

import (
	"context"
	"log"

	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain"
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

// SendNotification ...
func (s Service) SendNotification(ctx context.Context, notif domain.Notification, ch chan []*domain.User) (chan bool, chan error) {
	done := make(chan bool)
	errCh := make(chan error)

	go func() {
		for bulk := range ch {
			log.Printf("Sending %d Emails via Sendgrid", len(bulk))
		}

		done <- true
		close(done)
		close(errCh)
	}()

	return done, errCh
}
