package sendgrid

import (
	"context"
	"log"

	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain"
)

// SenderInput ...
type SenderInput struct {
}

// Sender ...
type Sender struct {
	in SenderInput
}

// NewSender ...
func NewSender(in SenderInput) (*Sender, error) {
	return &Sender{in: in}, nil
}

// SendNotification ...
func (s Sender) SendNotification(ctx context.Context, notif domain.Notification, ch chan []*domain.User) (chan bool, chan error) {
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
