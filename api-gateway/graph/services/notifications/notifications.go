package notifications

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"

	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/model"

	pb "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc/proto"
)

// ServiceInput ...
type ServiceInput struct {
	Tracer         trace.Tracer
	FrontendClient pb.NotificationsServiceFrontendClient
	WorkerClient   pb.NotificationsServiceWorkerClient
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

	if in.FrontendClient == nil {
		return nil, fmt.Errorf("Missing required dependency: FrontendClient")
	}

	if in.WorkerClient == nil {
		return nil, fmt.Errorf("Missing required dependency: WorkerClient")
	}

	return &Service{in: in}, nil
}

// ListNotifications ...
func (s Service) ListNotifications(ctx context.Context) ([]*model.Notification, error) {
	ctx, span := s.in.Tracer.Start(ctx, "graph.services.notifications.ListNotifications")
	defer span.End()

	res, err := s.in.FrontendClient.ListNotifications(ctx, &pb.ListNotificationsRequest{})
	if err != nil {
		return nil, err
	}

	notifications := []*model.Notification{}
	for _, notif := range res.Notifications {
		notifications = append(notifications, translate(notif))
	}

	return notifications, nil
}

// GetNotification ...
func (s Service) GetNotification(ctx context.Context, id string) (*model.Notification, error) {
	ctx, span := s.in.Tracer.Start(ctx, "graph.services.notifications.GetNotification")
	defer span.End()

	res, err := s.in.FrontendClient.GetNotification(ctx, &pb.GetNotificationRequest{
		NotificationId: id,
	})
	if err != nil {
		return nil, err
	}

	return translate(res.Notification), nil
}

// SendNotification ...
func (s Service) SendNotification(ctx context.Context, id string) error {
	ctx, span := s.in.Tracer.Start(ctx, "graph.services.notifications.SendNotification")
	defer span.End()

	_, err := s.in.WorkerClient.SendNotification(ctx, &pb.SendNotificationRequest{
		NotificationId: id,
	})

	return err
}

var dtoToModelChannel = map[pb.NotificationChannel]model.NotificationChannel{
	pb.NotificationChannel_EMAIL:   model.NotificationChannelEmail,
	pb.NotificationChannel_SMS:     model.NotificationChannelSms,
	pb.NotificationChannel_BROWSER: model.NotificationChannelBrowser,
}

func translate(dto *pb.Notification) *model.Notification {
	return &model.Notification{
		ID:      dto.Id,
		AppKey:  dto.AppKey,
		Name:    dto.Name,
		Channel: dtoToModelChannel[dto.Channel],
		Segment: &model.Segment{
			ID: dto.SegmentId,
		},
	}
}
