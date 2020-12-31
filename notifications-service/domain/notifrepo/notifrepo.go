package notifrepo

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain"
)

// RepositoryInput ...
type RepositoryInput struct {
}

// Repository ...
type Repository struct {
}

// NewRepository ...
func NewRepository(in RepositoryInput) (*Repository, error) {
	return &Repository{}, nil
}

// GetNotification ...
func (r Repository) GetNotification(ctx context.Context, id string) (*domain.Notification, error) {
	log.Printf("Fetch notification %s from database", id)

	notifications, err := r.ListNotifications(ctx)
	if err != nil {
		return nil, err
	}

	for _, notification := range notifications {
		if notification.ID == id {
			return notification, nil
		}
	}

	return nil, domain.ErrEntityNotFound
}

// ListNotifications ...
func (r Repository) ListNotifications(ctx context.Context) ([]*domain.Notification, error) {
	log.Printf("Fetch notifications from database")

	file, err := os.Open("config/notifications.json")
	if err != nil {
		return nil, err
	}

	notifications := []*domain.Notification{}

	parser := json.NewDecoder(file)
	if err := parser.Decode(&notifications); err != nil {
		return nil, err
	}

	return notifications, nil
}
