package notifications

import (
	"context"
	"fmt"

	pb "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc/proto"
	"google.golang.org/grpc"

	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/model"
	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/services"
)

// ServiceInput ...
type ServiceInput struct {
	ServerAddress string
}

// Service ...
type Service struct {
	in ServiceInput

	cli pb.NotificationsServiceFrontendClient
}

// NewService ...
func NewService(in ServiceInput) (*Service, error) {
	if in.ServerAddress == "" {
		return nil, fmt.Errorf("Missing required dependency: ServerAddress")
	}

	return &Service{in: in}, nil
}

// Connect ...
func (s *Service) Connect() error {
	conn, err := grpc.Dial(s.in.ServerAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}

	s.cli = pb.NewNotificationsServiceFrontendClient(conn)

	return nil
}

// ListNotifications ...
func (s Service) ListNotifications(ctx context.Context) ([]*model.Notification, error) {
	if s.cli == nil {
		return nil, fmt.Errorf("client not connected to NotificationsService gRPC server")
	}

	res, err := s.cli.ListNotifications(ctx, &pb.ListNotificationsRequest{})
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
	if s.cli == nil {
		return nil, fmt.Errorf("client not connected to NotificationsService gRPC server")
	}

	res, err := s.cli.GetNotification(ctx, &pb.GetNotificationRequest{
		NotificationId: id,
	})
	if err != nil {
		return nil, err
	}

	return translate(res.Notification), nil
}

// SendNotification ...
func (s Service) SendNotification(ctx context.Context, id string) error {
	if s.cli == nil {
		return fmt.Errorf("client not connected to NotificationsService gRPC server")
	}

	return services.ErrNotImplemented
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
