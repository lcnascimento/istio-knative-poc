package grpc

import (
	"context"
	"fmt"
	"log"

	pb "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc/proto"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain"
)

// WorkerInput ...
type WorkerInput struct {
	Sender domain.NotificationsSender
}

// Worker ...
type Worker struct {
	in WorkerInput

	pb.UnimplementedNotificationsServiceWorkerServer
}

// NewWorker ...
func NewWorker(in WorkerInput) (*Worker, error) {
	if in.Sender == nil {
		return nil, fmt.Errorf("Missing NotificationsSender dependency")
	}

	return &Worker{in: in}, nil
}

// SendNotification ...
func (s Worker) SendNotification(ctx context.Context, in *pb.SendNotificationRequest) (*wrapperspb.BoolValue, error) {
	if err := s.in.Sender.SendNotification(ctx, in.NotificationId); err != nil {
		log.Printf("Error sending notification %s: %s", in.NotificationId, err.Error())
		return wrapperspb.Bool(false), err
	}

	return wrapperspb.Bool(true), nil
}
