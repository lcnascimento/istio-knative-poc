package grpc

import (
	"context"
	"fmt"

	pb "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc/proto"
)

const (
	address       = "localhost:8084"
	numberOfBulks = 3
)

// ErrNotImplemented ...
var ErrNotImplemented error = fmt.Errorf("method not implemented yet")

// Server ...
type Server struct {
	pb.UnimplementedNotificationsServiceServer
}

// NewServer ...
func NewServer() *Server {
	return &Server{}
}

// GetNotification ...
func (s Server) GetNotification(ctx context.Context, in *pb.GetNotificationRequest) (*pb.GetNotificationResponse, error) {
	return &pb.GetNotificationResponse{
		Notification: &pb.Notification{
			Id:      "1",
			Name:    "Notification 1",
			AppKey:  "AppKey 1",
			Channel: "email",
		},
	}, nil
}

// ListNotifications ...
func (s Server) ListNotifications(ctx context.Context, in *pb.ListNotificationsRequest) (*pb.ListNotificationsResponse, error) {
	return &pb.ListNotificationsResponse{
		Data: []*pb.Notification{
			{
				Id:      "1",
				Name:    "Notification 1",
				AppKey:  "AppKey 1",
				Channel: "email",
			},
			{
				Id:      "2",
				Name:    "Notification 2",
				AppKey:  "AppKey 2",
				Channel: "email",
			},
			{
				Id:      "3",
				Name:    "Notification 3",
				AppKey:  "AppKey 3",
				Channel: "email",
			},
		},
	}, nil
}

// SendNotification ...
func (s Server) SendNotification(ctx context.Context, in *pb.SendNotificationRequest) (*pb.Void, error) {
	return nil, ErrNotImplemented
}
