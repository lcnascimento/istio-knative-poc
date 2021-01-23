package grpc

import (
	"context"
	"fmt"
	"log"

	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/trace"

	pb "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc/proto"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain"
)

// FrontendInput ...
type FrontendInput struct {
	Tracer   trace.Tracer
	Repo     domain.NotificationsRepository
	Enqueuer domain.NotificationsEnqueuer
}

// Frontend ...
type Frontend struct {
	in FrontendInput

	pb.UnimplementedNotificationsServiceFrontendServer
}

// NewFrontend ...
func NewFrontend(in FrontendInput) (*Frontend, error) {
	if in.Tracer == nil {
		return nil, fmt.Errorf("Missing required dependency: Tracer")
	}

	if in.Repo == nil {
		return nil, fmt.Errorf("Missing required dependency: Repo")
	}

	if in.Enqueuer == nil {
		return nil, fmt.Errorf("Missing required dependency: Enqueuer")
	}

	return &Frontend{in: in}, nil
}

// GetNotification ...
func (s Frontend) GetNotification(ctx context.Context, in *pb.GetNotificationRequest) (*pb.GetNotificationResponse, error) {
	ctx, span := s.in.Tracer.Start(ctx, "application.grpc.frontend.GetNotification")
	defer span.End()

	span.SetAttributes(label.String("notification_id", in.NotificationId))

	notif, err := s.in.Repo.GetNotification(ctx, in.NotificationId)
	if err != nil {
		log.Printf("could not get notification %s: %s", in.NotificationId, err.Error())
		return nil, err
	}

	return &pb.GetNotificationResponse{Notification: notif.ToGRPCDTO()}, nil
}

// ListNotifications ...
func (s Frontend) ListNotifications(ctx context.Context, in *pb.ListNotificationsRequest) (*pb.ListNotificationsResponse, error) {
	ctx, span := s.in.Tracer.Start(ctx, "application.grpc.frontend.ListNotifications")
	defer span.End()

	notifs, err := s.in.Repo.ListNotifications(ctx)
	if err != nil {
		log.Printf("could not list notifications: %s", err.Error())
		return nil, err
	}

	response := []*pb.Notification{}
	for _, notif := range notifs {
		response = append(response, notif.ToGRPCDTO())
	}

	return &pb.ListNotificationsResponse{Notifications: response}, nil
}

// EnqueueSendingNotification ...
func (s Frontend) EnqueueSendingNotification(ctx context.Context, in *pb.SendNotificationRequest) (*wrapperspb.BoolValue, error) {
	ctx, span := s.in.Tracer.Start(ctx, "application.grpc.frontend.EnqueueSendingNotification")
	defer span.End()

	span.SetAttributes(label.String("notification_id", in.NotificationId))

	if err := s.in.Enqueuer.EnqueueNotification(ctx, in.NotificationId); err != nil {
		log.Printf("could not enqueue notification %s: %s", in.NotificationId, err.Error())
		return wrapperspb.Bool(false), err
	}

	return wrapperspb.Bool(true), nil
}
