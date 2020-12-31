package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/lcnascimento/istio-knative-poc/notifications-service/application/grpc/proto"
)

const address = "localhost:8084"

func main() {
	ctx := context.Background()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	cli := pb.NewNotificationsServiceClient(conn)

	res, err := cli.ListNotifications(ctx, &pb.ListNotificationsRequest{})
	if err != nil {
		log.Fatalf("could not do request %v", err)
	}

	log.Println(res.Data)
}
