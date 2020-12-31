package grpc

import (
	"fmt"
	"log"
	"sync"
	"time"

	pb "github.com/lcnascimento/istio-knative-poc/segments-service/application/grpc/proto"
)

const numberOfBulks = 3

// Server ...
type Server struct {
	in         ServerInput
	lastUserID int

	pb.UnimplementedSegmentsServiceServer
}

// ServerInput ...
type ServerInput struct {
	NetworkSimulationTimeout time.Duration
}

// NewServer ...
func NewServer(in ServerInput) *Server {
	return &Server{in: in}
}

// GetSegmentUsers ...
func (s Server) GetSegmentUsers(in *pb.GetSegmentUsersRequest, srv pb.SegmentsService_GetSegmentUsersServer) error {
	log.Println("Fetching users from segment ", in.SegmentId)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for i := 0; i < numberOfBulks; i++ {
			resp := pb.GetSegmentUsersResponse{
				Total: numberOfBulks * in.Size,
			}

			resp.Data = s.getUsers(in.Size)

			if err := srv.Send(&resp); err != nil {
				log.Println("could not send message to client: ", err)
			}

			log.Printf("Message %d sent to client", i+1)
		}

		log.Println("All messages sent to client")

		wg.Done()
	}()

	wg.Wait()

	return nil
}

func (s *Server) getUsers(numUsers int32) []*pb.User {
	log.Println("Fetching users from database")
	time.Sleep(s.in.NetworkSimulationTimeout)

	users := []*pb.User{}
	for i := int32(0); i < numUsers; i++ {
		s.lastUserID++
		users = append(users, &pb.User{Id: fmt.Sprintf("%d", s.lastUserID)})
	}

	return users
}
