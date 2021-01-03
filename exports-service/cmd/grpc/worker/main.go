package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/env"

	app "github.com/lcnascimento/istio-knative-poc/exports-service/application/grpc"
	pb "github.com/lcnascimento/istio-knative-poc/exports-service/application/grpc/proto"

	"github.com/lcnascimento/istio-knative-poc/exports-service/domain/exportator"
	repo "github.com/lcnascimento/istio-knative-poc/exports-service/domain/repository"
	"github.com/lcnascimento/istio-knative-poc/exports-service/domain/segments"
)

func main() {
	repo, err := repo.NewService(repo.ServiceInput{})
	if err != nil {
		log.Fatalf("can not initialize ExportsRepository: %v", err)
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
		log.Fatalf("can not initialize SegmentsService: %v", err)
	}

	if err := segments.Connect(); err != nil {
		log.Fatalf("can not connect to Segments gRPC server: %v", err)
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

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", env.MustGetInt("PORT")))
	if err != nil {
		log.Fatalf("can not initialize gRPC server: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterExportsServiceWorkerServer(s, worker)

	log.Println("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("could not initialize grpc server: %v", err)
	}
}
