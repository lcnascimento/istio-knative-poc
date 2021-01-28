package main

import (
	"context"
	"fmt"

	ggrpc "google.golang.org/grpc"

	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/env"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/errors"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/grpc"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/log"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/metrics"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/tracing"

	app "github.com/lcnascimento/istio-knative-poc/exports-service/application/grpc"
	pb "github.com/lcnascimento/istio-knative-poc/exports-service/application/grpc/proto"
	segmentsPb "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc/proto"

	"github.com/lcnascimento/istio-knative-poc/exports-service/domain/exportator"
	"github.com/lcnascimento/istio-knative-poc/exports-service/domain/repository"
	"github.com/lcnascimento/istio-knative-poc/exports-service/domain/segments"
)

const applicationName = "exports-service-worker"

func main() {
	ctx := context.Background()

	log, err := log.NewClient(log.ClientInput{Level: log.DebugLevel})

	tracer, flush, err := tracing.Init(tracing.TracerInput{
		AgentEndpoint:   fmt.Sprintf("%s:%d", env.MustGetString("JAEGER_AGENT_HOST"), env.MustGetInt("JAEGER_AGENT_PORT")),
		ApplicationName: applicationName,
		TracerName:      fmt.Sprintf("%s-tracer", applicationName),
	})
	if err != nil {
		log.Critical(ctx, errors.New(fmt.Sprintf("can not initialize Tracer %s", err.Error())))
		return
	}
	defer flush()

	_, err = metrics.Init(metrics.MeterInput{
		ApplicationName: applicationName,
		ServerPort:      env.MustGetInt("PROMETHEUS_METRICS_EXPORTER_PORT"),
	})
	if err != nil {
		log.Critical(ctx, errors.New(fmt.Sprintf("can not initialize Meter %s", err.Error())))
		return
	}

	repository, err := repository.NewService(repository.ServiceInput{
		Tracer: tracer,
		Logger: log,
	})
	if err != nil {
		log.Critical(ctx, errors.New(fmt.Sprintf("can not initialize Exportsrepositorysitory: %v", err)))
		return
	}

	segmentsGRPCClient, err := grpc.NewClient(grpc.ClientInput{
		ServerAddress: fmt.Sprintf(
			"%s:%d",
			env.MustGetString("SEGMENTS_SERVICE_SERVER_HOST"),
			env.MustGetInt("SEGMENTS_SERVICE_SERVER_PORT"),
		),
		Tracer: tracer,
		Logger: log,
	})
	if err != nil {
		log.Critical(ctx, errors.New(fmt.Sprintf("can not initialize SegmentsService's gRPC client: %s", err.Error())))
		return
	}

	segmentsGRPCClientConn, err := segmentsGRPCClient.Connect(ctx)
	if err != nil {
		log.Critical(ctx, errors.New(fmt.Sprintf("can not connect to SegmentsService's gRPC client: %s", err.Error())))
		return
	}

	segments, err := segments.NewService(segments.ServiceInput{
		BulkSize: env.MustGetInt("SEGMENTS_SERVICE_BULK_SIZE"),
		Tracer:   tracer,
		Logger:   log,
		Client:   segmentsPb.NewSegmentsServiceFrontendClient(segmentsGRPCClientConn),
	})
	if err != nil {
		log.Critical(ctx, errors.New(fmt.Sprintf("can not initialize SegmentsService %s", err.Error())))
		return
	}

	exportator, err := exportator.NewService(exportator.ServiceInput{
		Tracer:   tracer,
		Logger:   log,
		Repo:     repository,
		Segments: segments,
	})
	if err != nil {
		log.Critical(ctx, errors.New(fmt.Sprintf("can not initialize ExportsEnqueuer: %v", err)))
		return
	}

	worker, err := app.NewWorker(app.WorkerInput{
		Tracer:     tracer,
		Exportator: exportator,
	})
	if err != nil {
		log.Critical(ctx, errors.New(fmt.Sprintf("can not initialize server: %v", err)))
		return
	}

	s, err := grpc.NewServer(grpc.ServerInput{
		Port:   env.MustGetInt("PORT"),
		Tracer: tracer,
		Logger: log,
		Registrator: func(srv ggrpc.ServiceRegistrar) {
			pb.RegisterExportsServiceWorkerServer(srv, worker)
		},
	})
	if err != nil {
		log.Critical(ctx, errors.New(fmt.Sprintf("can not create gRPC server %s", err.Error())))
		return
	}

	if err := s.Listen(ctx); err != nil {
		log.Critical(ctx, errors.New(fmt.Sprintf("could not initialize grpc server: %s", err.Error())))
		return
	}
}
