package main

import (
	"context"
	"fmt"
	"time"

	ggrpc "google.golang.org/grpc"

	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/env"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/errors"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/grpc"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/log"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/tracing"

	app "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc"
	pb "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc/proto"

	"github.com/lcnascimento/istio-knative-poc/segments-service/domain/repository"
)

func main() {
	ctx := context.Background()

	log, err := log.NewClient(log.ClientInput{Level: log.DebugLevel})

	tracer, flush, err := tracing.Init(tracing.TracerInput{
		AgentEndpoint: fmt.Sprintf("%s:%d", env.MustGetString("JAEGER_AGENT_HOST"), env.MustGetInt("JAEGER_AGENT_PORT")),
		ServiceName:   "segments-service-frontend",
		TracerName:    "segments-service-frontend-tracer",
	})
	if err != nil {
		log.Critical(ctx, errors.New(fmt.Sprintf("can not initialize Tracer %s", err.Error())))
		return
	}
	defer flush()

	repo, err := repository.NewService(repository.ServiceInput{
		NetworkDelay:            time.Millisecond * time.Duration(env.MustGetInt("NETWORK_DELAY_IN_MILLISECONDS")),
		NumberOfUsersInSegments: env.MustGetInt("NUMBER_OF_MOCKED_USERS_IN_SEGMENTS"),
		Tracer:                  tracer,
		Logger:                  log,
	})
	if err != nil {
		log.Critical(ctx, errors.New(fmt.Sprintf("could not initialize SegmentsRepository: %s", err.Error())))
		return
	}

	frontend, err := app.NewFrontend(app.FrontendInput{
		Tracer: tracer,
		Logger: log,
		Repo:   repo,
	})
	if err != nil {
		log.Critical(ctx, errors.New(fmt.Sprintf("could not initialize Frontend: %v", err.Error())))
		return
	}

	s, err := grpc.NewServer(grpc.ServerInput{
		Port:   env.MustGetInt("PORT"),
		Tracer: tracer,
		Logger: log,
		Registrator: func(srv ggrpc.ServiceRegistrar) {
			pb.RegisterSegmentsServiceFrontendServer(srv, frontend)
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
