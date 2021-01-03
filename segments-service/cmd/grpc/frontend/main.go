package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/env"

	app "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc"
	pb "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc/proto"

	"github.com/lcnascimento/istio-knative-poc/segments-service/domain/repository"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", env.MustGetInt("PORT")))
	if err != nil {
		log.Fatalf("can not initialize server %v", err)
	}

	repo, err := repository.NewService(repository.ServiceInput{
		NetworkDelay:            time.Millisecond * time.Duration(env.MustGetInt("NETWORK_DELAY_IN_MILLISECONDS")),
		NumberOfUsersInSegments: env.MustGetInt("NUMBER_OF_MOCKED_USERS_IN_SEGMENTS"),
	})
	if err != nil {
		log.Fatalf("could not initialize SegmentsRepository: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSegmentsServiceFrontendServer(s, app.NewFrontend(app.FrontendInput{Repo: repo}))

	log.Println("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("could not initialize grpc server: %v", err)
	}
}
