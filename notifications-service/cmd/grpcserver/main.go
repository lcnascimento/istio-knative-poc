package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	app "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc"
	pb "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc/proto"
)

const address = "localhost:8084"

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("can not initialize server %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNotificationsServiceServer(s, app.NewServer())

	log.Println("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("could not initialize grpc server: %v", err)
	}
}
