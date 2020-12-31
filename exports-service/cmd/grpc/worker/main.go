package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	app "github.com/lcnascimento/istio-knative-poc/exports-service/application/grpc"
	pb "github.com/lcnascimento/istio-knative-poc/exports-service/application/grpc/proto"

	"github.com/lcnascimento/istio-knative-poc/exports-service/domain/exportator"
	repo "github.com/lcnascimento/istio-knative-poc/exports-service/domain/repository"
	"github.com/lcnascimento/istio-knative-poc/exports-service/domain/segments"
)

const address = "localhost:8086"

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("can not initialize gRPC server: %v", err)
	}

	repo, err := repo.NewService(repo.ServiceInput{})
	if err != nil {
		log.Fatalf("can not initialize ExportsRepository: %v", err)
	}

	segments, err := segments.NewService(segments.ServiceInput{
		ServerAddress: "localhost:8083",
	})
	if err != nil {
		log.Fatalf("can not initialize SegmentsService: %v", err)
	}

	if err := segments.Connect(); err != nil {
		log.Fatalf("can not connect to Segments gRPX server: %v", err)
	}

	exportator, err := exportator.NewService(exportator.ServiceInput{
		Repo:     repo,
		Segments: segments,
	})
	if err != nil {
		log.Fatalf("can not initialize ExportsEnqueuer: %v", err)
	}

	worker, err := app.NewWorker(app.WorkerInput{
		Exportator: exportator,
	})
	if err != nil {
		log.Fatalf("can not initialize server: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterExportsServiceWorkerServer(s, worker)

	log.Println("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("could not initialize grpc server: %v", err)
	}
}
