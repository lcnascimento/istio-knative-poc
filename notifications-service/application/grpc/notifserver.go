package grpc

import (
	"context"
	"fmt"
	"log"

	pb "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc/proto"
	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain"
)

// ServerInput ...
type ServerInput struct {
	Repo     domain.NotificationsRepository
	Enqueuer domain.NotificationsEnqueuer
}

// Server ...
type Server struct {
	in ServerInput

	pb.UnimplementedNotificationsServiceServer
}

// NewServer ...
func NewServer(in ServerInput) (*Server, error) {
	if in.Repo == nil {
		return nil, fmt.Errorf("Missing NotificationsRepository dependency")
	}

	if in.Enqueuer == nil {
		return nil, fmt.Errorf("Missing NotificationsEnqueuer dependency")
	}

	return &Server{in: in}, nil
}

// GetNotification ...
func (s Server) GetNotification(ctx context.Context, in *pb.GetNotificationRequest) (*pb.GetNotificationResponse, error) {
	notif, err := s.in.Repo.GetNotification(ctx, in.NotificationId)
	if err != nil {
		log.Printf("could not get notification %s: %s", in.NotificationId, err.Error())
		return nil, err
	}

	return &pb.GetNotificationResponse{Notification: notif.ToGRPCDTO()}, nil
}

// ListNotifications ...
func (s Server) ListNotifications(ctx context.Context, in *pb.ListNotificationsRequest) (*pb.ListNotificationsResponse, error) {
	notifs, err := s.in.Repo.ListNotifications(ctx)
	if err != nil {
		log.Printf("could not list notifications: %s", err.Error())
		return nil, err
	}

	response := []*pb.Notification{}
	for _, notif := range notifs {
		response = append(response, notif.ToGRPCDTO())
	}

	return &pb.ListNotificationsResponse{Data: response}, nil
}

// EnqueueSendingNotification ...
func (s Server) EnqueueSendingNotification(ctx context.Context, in *pb.SendNotificationRequest) (*pb.Void, error) {
	if err := s.in.Enqueuer.EnqueueNotification(ctx, in.NotificationId); err != nil {
		log.Printf("could not enqueue notification %s: %s", in.NotificationId, err.Error())
		return nil, err
	}

	return &pb.Void{}, nil
}
