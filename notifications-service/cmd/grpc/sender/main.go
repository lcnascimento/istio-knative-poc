package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	app "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc"
	pb "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc/proto"

	repo "github.com/lcnascimento/istio-knative-poc/notifications-service/domain/notifrepo"
	sender "github.com/lcnascimento/istio-knative-poc/notifications-service/domain/notifsender"

	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain"
	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain/firebase"
	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain/movile"
	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain/segments"
	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain/sendgrid"
)

const address = "localhost:8084"

func main() {
	repo, err := repo.NewRepository(repo.RepositoryInput{})
	if err != nil {
		log.Fatalf("can not initialize NotificationsRepository %v", err)
	}

	segments, err := segments.NewService(segments.ServiceInput{
		ServerAddress: "localhost:8083",
	})
	if err != nil {
		log.Fatalf("can not initialize SegmentsService %v", err)
	}

	if err := segments.Connect(); err != nil {
		log.Fatalf("can not initialize SegmentsService gRPC Server %v", err)
	}

	movile, err := movile.NewSender(movile.SenderInput{})
	if err != nil {
		log.Fatalf("can not initialize MovileService %v", err)
	}

	sendgrid, err := sendgrid.NewSender(sendgrid.SenderInput{})
	if err != nil {
		log.Fatalf("can not initialize SendgridService %v", err)
	}

	firebase, err := firebase.NewSender(firebase.SenderInput{})
	if err != nil {
		log.Fatalf("can not initialize FirebaseService %v", err)
	}

	providers := map[domain.Channel]domain.NotificationProvider{
		domain.EmailChannel:       sendgrid,
		domain.SMSChannel:         movile,
		domain.BrowserPushChannel: firebase,
	}

	sender, err := sender.NewSender(sender.SenderInput{
		Repo:      repo,
		Segments:  segments,
		Providers: providers,
	})
	if err != nil {
		log.Fatalf("can not initialize NotificationsSender %v", err)
	}

	server, err := app.NewSender(app.SenderInput{Sender: sender})
	if err != nil {
		log.Fatalf("can not initialize GRPCNotificationsSender %v", err)
	}

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("can not initialize gRPC server %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNotificationsSenderServiceServer(s, server)

	log.Println("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("could not initialize grpc server: %v", err)
	}
}
