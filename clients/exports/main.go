package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/lcnascimento/istio-knative-poc/exports-service/application/grpc/proto"
)

const address = "localhost:8086"

// func main() {
// 	ctx := context.Background()

// 	conn, err := grpc.Dial(address, grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("can not connect with server %v", err)
// 	}

// 	cli := pb.NewExportsServiceFrontendClient(conn)

// 	res, err := cli.EnqueueExport(ctx, &pb.EnqueueExportRequest{
// 		Id: "6",
// 	})
// 	if err != nil {
// 		log.Fatalf("could not do request %v", err)
// 	}

// 	log.Println(res)
// }

func main() {
	ctx := context.Background()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	cli := pb.NewExportsServiceWorkerClient(conn)

	cli.ProcessExport(ctx, &pb.ProcessExportRequest{ExportId: "1"})
	if err != nil {
		log.Fatalf("could not do request %v", err)
	}

	log.Println("Operation done successfuly")
}
