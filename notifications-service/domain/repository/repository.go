package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

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

// GetNotification ...
func (r Service) GetNotification(ctx context.Context, id string) (*domain.Notification, error) {
	ctx, span := r.in.Tracer.Start(ctx, "domain.repository.GetNotification")
	defer span.End()

	r.in.Logger.Info(ctx, "Fetch notification %s from database", id)

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
func (r Service) ListNotifications(ctx context.Context) ([]*domain.Notification, error) {
	ctx, span := r.in.Tracer.Start(ctx, "domain.repository.ListNotifications")
	defer span.End()

	r.in.Logger.Info(ctx, "Fetch notifications from database")

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
