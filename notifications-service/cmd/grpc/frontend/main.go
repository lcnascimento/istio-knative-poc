package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	app "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc"
	pb "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc/proto"

	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain/enqueuer"
	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain/repository"
)

const address = "localhost:8084"

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("can not initialize gRPC server %v", err)
	}

	enqueuer, err := enqueuer.NewService(enqueuer.ServiceInput{})
	if err != nil {
		log.Fatalf("can not initialize NotificationsEnqueuer %v", err)
	}

	repo, err := repository.NewService(repository.ServiceInput{})
	if err != nil {
		log.Fatalf("can not initialize NotificationsRepository %v", err)
	}

	frontend, err := app.NewFrontend(app.FrontendInput{
		Repo:     repo,
		Enqueuer: enqueuer,
	})
	if err != nil {
		log.Fatalf("can not initialize server %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNotificationsServiceFrontendServer(s, frontend)

	log.Println("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("could not initialize grpc server: %v", err)
	}
}
