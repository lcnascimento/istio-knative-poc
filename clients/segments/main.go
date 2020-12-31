package main

import (
	"context"
	"log"
	"sync"

	"google.golang.org/grpc"

	pb "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc/proto"
)

const address = "localhost:8083"

func main() {
	ctx := context.Background()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	cli := pb.NewSegmentsServiceClient(conn)

	stream, err := cli.GetSegmentUsers(ctx, &pb.GetSegmentUsersRequest{
		SegmentId: "1",
		Size:      1,
	})
	if err != nil {
		log.Fatalf("could not create GetSegmentUsers stream %v", err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				wg.Done()
				return
			}

			log.Println(resp.Data)
		}
	}()

	wg.Wait()
	log.Println("Finished")
}
