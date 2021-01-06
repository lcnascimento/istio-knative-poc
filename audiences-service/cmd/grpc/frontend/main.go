package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/env"

	app "github.com/lcnascimento/istio-knative-poc/audiences-service/application/grpc"
	pb "github.com/lcnascimento/istio-knative-poc/audiences-service/application/grpc/proto"

	enqueuer "github.com/lcnascimento/istio-knative-poc/audiences-service/domain/enqueuer"
	repo "github.com/lcnascimento/istio-knative-poc/audiences-service/domain/repository"
)

func main() {
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

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", env.MustGetInt("PORT")))
	if err != nil {
		log.Fatalf("can not initialize gRPC server %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAudiencesServiceFrontendServer(s, frontend)

	log.Println("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("could not initialize grpc server: %v", err)
	}
}
