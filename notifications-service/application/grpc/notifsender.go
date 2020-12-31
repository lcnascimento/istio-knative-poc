package grpc

import (
	"context"
	"fmt"
	"log"

	pb "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc/proto"
	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain"
)

// SenderInput ...
type SenderInput struct {
	Sender domain.NotificationsSender
}

// Sender ...
type Sender struct {
	in SenderInput

	pb.UnimplementedNotificationsSenderServiceServer
}

// NewSender ...
func NewSender(in SenderInput) (*Sender, error) {
	if in.Sender == nil {
		return nil, fmt.Errorf("Missing NotificationsSender dependency")
	}

	return &Sender{in: in}, nil
}

// SendNotification ...
func (s Sender) SendNotification(ctx context.Context, in *pb.SendNotificationRequest) (*pb.Void, error) {
	if err := s.in.Sender.SendNotification(ctx, in.NotificationId); err != nil {
		log.Printf("Error sending notification %s: %s", in.NotificationId, err.Error())
		return nil, err
	}

	return &pb.Void{}, nil
}
