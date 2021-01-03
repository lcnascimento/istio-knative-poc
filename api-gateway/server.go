package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/lcnascimento/istio-knative-poc/go-libs/infra/env"

	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph"
	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/generated"

	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/services/exports"
	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/services/notifications"
	"github.com/lcnascimento/istio-knative-poc/api-gateway/graph/services/segments"
)

func main() {
	exports, err := exports.NewService(exports.ServiceInput{})
	if err != nil {
		log.Fatalf("could not initialize ExportService: %s", err.Error())
	}

	notificationsAddress := fmt.Sprintf(
		"%s:%d",
		env.MustGetString("NOTIFICATIONS_SERVICE_SERVER_HOST"),
		env.MustGetInt("NOTIFICATIONS_SERVICE_SERVER_PORT"),
	)
	notifications, err := notifications.NewService(notifications.ServiceInput{ServerAddress: notificationsAddress})
	if err != nil {
		log.Fatalf("could not initialize NotificationsService: %s", err.Error())
	}

	if err := notifications.Connect(); err != nil {
		log.Fatalf("could not connect to NotificationsService gRPC server: %s", err.Error())
	}

	segmentsAddress := fmt.Sprintf(
		"%s:%d",
		env.MustGetString("SEGMENTS_SERVICE_SERVER_HOST"),
		env.MustGetInt("SEGMENTS_SERVICE_SERVER_PORT"),
	)
	segments, err := segments.NewService(segments.ServiceInput{ServerAddress: segmentsAddress})
	if err != nil {
		log.Fatalf("could not initialize SegmentsService: %s", err.Error())
	}

	if err := segments.Connect(); err != nil {
		log.Fatalf("could not connect to SegmentsService gRPC server: %s", err.Error())
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		NotificationsService: notifications,
		ExportsService:       exports,
		SegmentsService:      segments,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", env.MustGetInt("PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", env.MustGetInt("PORT")), nil))
}
