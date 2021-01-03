package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/env"

	app "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc"
	pb "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc/proto"

	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain"
	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain/firebase"
	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain/movile"
	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain/repository"
	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain/segments"
	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain/sender"
	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain/sendgrid"
)

func main() {
	repo, err := repository.NewService(repository.ServiceInput{})
	if err != nil {
		log.Fatalf("can not initialize NotificationsRepository %v", err)
	}

	segmentsAddress := fmt.Sprintf(
		"%s:%d",
		env.MustGetString("SEGMENTS_SERVICE_SERVER_HOST"),
		env.MustGetInt("SEGMENTS_SERVICE_SERVER_PORT"),
	)
	segments, err := segments.NewService(segments.ServiceInput{
		ServerAddress: segmentsAddress,
		BulkSize:      env.MustGetInt("SEGMENTS_SERVICE_BULK_SIZE"),
	})
	if err != nil {
		log.Fatalf("can not initialize SegmentsService %v", err)
	}

	if err := segments.Connect(); err != nil {
		log.Fatalf("can not initialize SegmentsService gRPC Server %v", err)
	}

	movile, err := movile.NewService(movile.ServiceInput{})
	if err != nil {
		log.Fatalf("can not initialize MovileService %v", err)
	}

	sendgrid, err := sendgrid.NewService(sendgrid.ServiceInput{})
	if err != nil {
		log.Fatalf("can not initialize SendgridService %v", err)
	}

	firebase, err := firebase.NewService(firebase.ServiceInput{})
	if err != nil {
		log.Fatalf("can not initialize FirebaseService %v", err)
	}

	providers := map[domain.Channel]domain.NotificationProvider{
		domain.EmailChannel:       sendgrid,
		domain.SMSChannel:         movile,
		domain.BrowserPushChannel: firebase,
	}

	sender, err := sender.NewService(sender.ServiceInput{
		Repo:      repo,
		Segments:  segments,
		Providers: providers,
	})
	if err != nil {
		log.Fatalf("can not initialize NotificationsSender %v", err)
	}

	worker, err := app.NewWorker(app.WorkerInput{Sender: sender})
	if err != nil {
		log.Fatalf("can not initialize GRPCNotificationsSender %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", env.MustGetInt("PORT")))
	if err != nil {
		log.Fatalf("can not initialize gRPC server %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNotificationsServiceWorkerServer(s, worker)

	log.Println("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("could not initialize grpc server: %v", err)
	}
}
