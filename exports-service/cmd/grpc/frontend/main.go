package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	app "github.com/lcnascimento/istio-knative-poc/exports-service/application/grpc"
	pb "github.com/lcnascimento/istio-knative-poc/exports-service/application/grpc/proto"

	enqueuer "github.com/lcnascimento/istio-knative-poc/exports-service/domain/enqueuer"
	repo "github.com/lcnascimento/istio-knative-poc/exports-service/domain/repository"
)

const address = "localhost:8085"

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("can not initialize gRPC server %v", err)
	}

	enqueuer, err := enqueuer.NewService(enqueuer.ServiceInput{})
	if err != nil {
		log.Fatalf("can not initialize ExportsEnqueuer %v", err)
	}

	repo, err := repo.NewService(repo.ServiceInput{})
	if err != nil {
		log.Fatalf("can not initialize ExportsRepository %v", err)
	}

	frontend, err := app.NewFrontend(app.FrontendInput{
		Repo:     repo,
		Enqueuer: enqueuer,
	})
	if err != nil {
		log.Fatalf("can not initialize server %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterExportsServiceFrontendServer(s, frontend)

	log.Println("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("could not initialize grpc server: %v", err)
	}
}
