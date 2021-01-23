package grpc

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/trace"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain"

	pb "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc/proto"
)

// WorkerInput ...
type WorkerInput struct {
	Tracer trace.Tracer
	Sender domain.NotificationsSender
}

// Worker ...
type Worker struct {
	in WorkerInput

	pb.UnimplementedNotificationsServiceWorkerServer
}

// NewWorker ...
func NewWorker(in WorkerInput) (*Worker, error) {
	if in.Tracer == nil {
		return nil, fmt.Errorf("Missing required dependency: Tracer")
	}

	if in.Sender == nil {
		return nil, fmt.Errorf("Missing required dependency: Sender")
	}

	return &Worker{in: in}, nil
}

// SendNotification ...
func (s Worker) SendNotification(ctx context.Context, in *pb.SendNotificationRequest) (*wrapperspb.BoolValue, error) {
	ctx, span := s.in.Tracer.Start(ctx, "application.grpc.worker.SendNotification")
	defer span.End()

	span.SetAttributes(label.String("notification_id", in.NotificationId))

	if err := s.in.Sender.SendNotification(ctx, in.NotificationId); err != nil {
		return wrapperspb.Bool(false), err
	}

	return wrapperspb.Bool(true), nil
}
