package main

import (
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	app "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc"
	pb "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc/proto"
)

const address = "localhost:8083"

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("can not initialize server %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSegmentsServiceServer(s, app.NewServer(app.ServerInput{NetworkSimulationTimeout: time.Second}))

	log.Println("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("could not initialize grpc server: %v", err)
	}
}
