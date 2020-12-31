package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	app "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc"
	pb "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc/proto"

	enqueuer "github.com/lcnascimento/istio-knative-poc/notifications-service/domain/notifenqueuer"
	repo "github.com/lcnascimento/istio-knative-poc/notifications-service/domain/notifrepo"
)

const address = "localhost:8084"

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("can not initialize gRPC server %v", err)
	}

	enqueuer, err := enqueuer.NewEnqueuer(enqueuer.EnqueuerInput{})
	if err != nil {
		log.Fatalf("can not initialize NotificationsEnqueuer %v", err)
	}

	repo, err := repo.NewRepository(repo.RepositoryInput{})
	if err != nil {
		log.Fatalf("can not initialize NotificationsRepository %v", err)
	}

	server, err := app.NewServer(app.ServerInput{
		Repo:     repo,
		Enqueuer: enqueuer,
	})
	if err != nil {
		log.Fatalf("can not initialize server %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNotificationsServiceServer(s, server)

	log.Println("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("could not initialize grpc server: %v", err)
	}
}
