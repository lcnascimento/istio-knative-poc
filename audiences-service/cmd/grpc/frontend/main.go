package main

import (
	"context"
	"fmt"

	ggrpc "google.golang.org/grpc"

	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/env"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/grpc"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/log"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/tracing"

	app "github.com/lcnascimento/istio-knative-poc/audiences-service/application/grpc"
	pb "github.com/lcnascimento/istio-knative-poc/audiences-service/application/grpc/proto"

	enqueuer "github.com/lcnascimento/istio-knative-poc/audiences-service/domain/enqueuer"
	repo "github.com/lcnascimento/istio-knative-poc/audiences-service/domain/repository"
)

func main() {
	ctx := context.Background()

	log, err := log.NewClient(log.ClientInput{Level: log.DebugLevel})

	tracer, flush, err := tracing.Init(tracing.TracerInput{
		AgentEndpoint: fmt.Sprintf("%s:%d", env.MustGetString("JAEGER_AGENT_HOST"), env.MustGetInt("JAEGER_AGENT_PORT")),
		ServiceName:   "audiences-service-frontend",
		TracerName:    "audiences-service-frontend-tracer",
	})
	if err != nil {
		log.Critical(ctx, fmt.Errorf("can not initialize Tracer %s", err.Error()))
		return
	}
	defer flush()

	enqueuer, err := enqueuer.NewService(enqueuer.ServiceInput{})
	if err != nil {
		log.Critical(ctx, fmt.Errorf("can not initialize ExportsEnqueuer %v", err))
		return
	}

	repo, err := repo.NewService(repo.ServiceInput{
		Tracer: tracer,
		Logger: log,
	})
	if err != nil {
		log.Critical(ctx, fmt.Errorf("can not initialize ExportsRepository %v", err))
		return
	}

	frontend, err := app.NewFrontend(app.FrontendInput{
		Tracer:   tracer,
		Repo:     repo,
		Enqueuer: enqueuer,
	})
	if err != nil {
		log.Critical(ctx, fmt.Errorf("can not initialize server %v", err))
		return
	}

	s, err := grpc.NewServer(grpc.ServerInput{
		Port:   env.MustGetInt("PORT"),
		Tracer: tracer,
		Logger: log,
		Registrator: func(srv ggrpc.ServiceRegistrar) {
			pb.RegisterAudiencesServiceFrontendServer(srv, frontend)
		},
	})
	if err != nil {
		log.Critical(ctx, fmt.Errorf("can not create gRPC server %s", err.Error()))
		return
	}

	if err := s.Listen(ctx); err != nil {
		log.Critical(ctx, fmt.Errorf("could not initialize grpc server: %s", err.Error()))
		return
	}
}
