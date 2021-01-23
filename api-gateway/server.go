package main

import (
	"context"
	"fmt"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/env"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/errors"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/grpc"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/log"
	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/tracing"

	infra "github.com/lcnascimento/istio-knative-poc/go-libs/infra"

	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph"
	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/generated"

	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/services/audiences"
	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/services/exports"
	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/services/notifications"
	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/services/segments"

	audiencesPb "github.com/lcnascimento/istio-knative-poc/audiences-service/application/grpc/proto"
	exportsPb "github.com/lcnascimento/istio-knative-poc/exports-service/application/grpc/proto"
	notifPb "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc/proto"
	segmentsPb "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc/proto"
)

func main() {
	ctx := context.Background()

	log, err := log.NewClient(log.ClientInput{Level: log.DebugLevel})

	tracer, flush, err := tracing.Init(tracing.TracerInput{
		AgentEndpoint: fmt.Sprintf("%s:%d", env.MustGetString("JAEGER_AGENT_HOST"), env.MustGetInt("JAEGER_AGENT_PORT")),
		ServiceName:   "api-gateway",
		TracerName:    "api-gateway-tracer",
	})
	if err != nil {
		log.Critical(ctx, errors.New(fmt.Sprintf("can not initialize Tracer %s", err.Error())))
		return
	}
	defer flush()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		AudiencesService:     buildAudiencesService(ctx, tracer, log),
		ExportsService:       buildExportsService(ctx, tracer, log),
		NotificationsService: buildNotificationsService(ctx, tracer, log),
		SegmentsService:      buildSegmentsService(ctx, tracer, log),
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", otelhttp.NewHandler(srv, "GraphQL endpoint"))

	log.Info(ctx, "connect to http://localhost:%d/ for GraphQL playground", env.MustGetInt("PORT"))
	log.Critical(ctx, http.ListenAndServe(fmt.Sprintf(":%d", env.MustGetInt("PORT")), nil))
}

func buildAudiencesService(ctx context.Context, tracer trace.Tracer, log infra.LogProvider) *audiences.Service {
	audiencesGRPCClient, err := grpc.NewClient(grpc.ClientInput{
		ServerAddress: fmt.Sprintf(
			"%s:%d",
			env.MustGetString("AUDIENCES_SERVICE_SERVER_HOST"),
			env.MustGetInt("AUDIENCES_SERVICE_SERVER_PORT"),
		),
		Tracer: tracer,
		Logger: log,
	})
	if err != nil {
		msg := fmt.Sprintf("can not initialize AudiencesService's gRPC client: %s", err.Error())
		log.Critical(ctx, errors.New(msg))
		panic(msg)
	}

	audiencesGRPCClientConn, err := audiencesGRPCClient.Connect(ctx)
	if err != nil {
		msg := fmt.Sprintf("can not connect to AudiencesService's gRPC client: %s", err.Error())
		log.Critical(ctx, errors.New(msg))
		panic(msg)
	}

	audiences, err := audiences.NewService(audiences.ServiceInput{
		Tracer: tracer,
		Client: audiencesPb.NewAudiencesServiceFrontendClient(audiencesGRPCClientConn),
	})
	if err != nil {
		msg := fmt.Sprintf("can not initialize AudiencesService: %s", err.Error())
		log.Critical(ctx, errors.New(msg))
		panic(msg)
	}

	return audiences
}

func buildExportsService(ctx context.Context, tracer trace.Tracer, log infra.LogProvider) *exports.Service {
	exportsGRPCClient, err := grpc.NewClient(grpc.ClientInput{
		ServerAddress: fmt.Sprintf(
			"%s:%d",
			env.MustGetString("EXPORTS_SERVICE_SERVER_HOST"),
			env.MustGetInt("EXPORTS_SERVICE_SERVER_PORT"),
		),
		Tracer: tracer,
		Logger: log,
	})
	if err != nil {
		msg := fmt.Sprintf("can not initialize ExportsService's gRPC client: %s", err.Error())
		log.Critical(ctx, errors.New(msg))
		panic(msg)
	}

	exportsGRPCClientConn, err := exportsGRPCClient.Connect(ctx)
	if err != nil {
		msg := fmt.Sprintf("can not connect to ExportsService's gRPC client: %s", err.Error())
		log.Critical(ctx, errors.New(msg))
		panic(msg)
	}

	exports, err := exports.NewService(exports.ServiceInput{
		Tracer: tracer,
		Client: exportsPb.NewExportsServiceFrontendClient(exportsGRPCClientConn),
	})
	if err != nil {
		msg := fmt.Sprintf("can not initialize ExportsService: %s", err.Error())
		log.Critical(ctx, errors.New(msg))
		panic(msg)
	}

	return exports
}

func buildSegmentsService(ctx context.Context, tracer trace.Tracer, log infra.LogProvider) *segments.Service {
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
		msg := fmt.Sprintf("can not initialize SegmentsService's gRPC client: %s", err.Error())
		log.Critical(ctx, errors.New(msg))
		panic(msg)
	}

	segmentsGRPCClientConn, err := segmentsGRPCClient.Connect(ctx)
	if err != nil {
		msg := fmt.Sprintf("can not connect to SegmentsService's gRPC client: %s", err.Error())
		log.Critical(ctx, errors.New(msg))
		panic(msg)
	}

	segments, err := segments.NewService(segments.ServiceInput{
		Tracer: tracer,
		Client: segmentsPb.NewSegmentsServiceFrontendClient(segmentsGRPCClientConn),
	})
	if err != nil {
		msg := fmt.Sprintf("can not initialize SegmentsService: %s", err.Error())
		log.Critical(ctx, errors.New(msg))
		panic(msg)
	}

	return segments
}

func buildNotificationsService(ctx context.Context, tracer trace.Tracer, log infra.LogProvider) *notifications.Service {
	notifGRPCClient, err := grpc.NewClient(grpc.ClientInput{
		ServerAddress: fmt.Sprintf(
			"%s:%d",
			env.MustGetString("NOTIFICATIONS_SERVICE_SERVER_HOST"),
			env.MustGetInt("NOTIFICATIONS_SERVICE_SERVER_PORT"),
		),
		Tracer: tracer,
		Logger: log,
	})
	if err != nil {
		msg := fmt.Sprintf("can not initialize NotificationsService's gRPC client: %s", err.Error())
		log.Critical(ctx, errors.New(msg))
		panic(msg)
	}

	notifGRPCClientConn, err := notifGRPCClient.Connect(ctx)
	if err != nil {
		msg := fmt.Sprintf("can not connect to NotificationsService's gRPC client: %s", err.Error())
		log.Critical(ctx, errors.New(msg))
		panic(msg)
	}

	notifications, err := notifications.NewService(notifications.ServiceInput{
		Tracer: tracer,
		Client: notifPb.NewNotificationsServiceFrontendClient(notifGRPCClientConn),
	})
	if err != nil {
		msg := fmt.Sprintf("can not initialize NotificationsService: %s", err.Error())
		log.Critical(ctx, errors.New(msg))
		panic(msg)
	}

	return notifications
}
