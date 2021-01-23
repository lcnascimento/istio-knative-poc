package main

import (
	"context"
	"fmt"

	ggrpc "google.golang.org/grpc"

	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/env"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/errors"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/grpc"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/log"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/tracing"

	app "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc"
	pb "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc/proto"

	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain/enqueuer"
	"github.com/lcnascimento/istio-knative-poc/notifications-service/domain/repository"
)

func main() {
	ctx := context.Background()

	log, err := log.NewClient(log.ClientInput{Level: log.DebugLevel})

	tracer, flush, err := tracing.Init(tracing.TracerInput{
		AgentEndpoint: fmt.Sprintf("%s:%d", env.MustGetString("JAEGER_AGENT_HOST"), env.MustGetInt("JAEGER_AGENT_PORT")),
		ServiceName:   "notifications-service-frontend",
		TracerName:    "notifications-service-frontend-tracer",
	})
	if err != nil {
		log.Critical(ctx, errors.New(fmt.Sprintf("can not initialize Tracer %s", err.Error())))
		return
	}
	defer flush()

	enqueuer, err := enqueuer.NewService(enqueuer.ServiceInput{})
	if err != nil {
		log.Critical(ctx, errors.New(fmt.Sprintf("can not initialize NotificationsEnqueuer %s", err.Error())))
		return
	}

	repo, err := repository.NewService(repository.ServiceInput{Logger: log, Tracer: tracer})
	if err != nil {
		log.Critical(ctx, errors.New(fmt.Sprintf("can not initialize NotificationsRepository %s", err.Error())))
		return
	}

	frontend, err := app.NewFrontend(app.FrontendInput{
		Tracer:   tracer,
		Repo:     repo,
		Enqueuer: enqueuer,
	})
	if err != nil {
		log.Critical(ctx, errors.New(fmt.Sprintf("can not initialize server %s", err.Error())))
		return
	}

	s, err := grpc.NewServer(grpc.ServerInput{
		Port:   env.MustGetInt("PORT"),
		Tracer: tracer,
		Logger: log,
		Registrator: func(srv ggrpc.ServiceRegistrar) {
			pb.RegisterNotificationsServiceFrontendServer(srv, frontend)
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
