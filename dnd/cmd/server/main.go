package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/rledford/monomicro-go/dnd/api/v1"
	randint "github.com/rledford/monomicro-go/randint/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

type server struct {
	pb.UnsafeDnDServiceServer
}

func (s *server) GetRoll(ctx context.Context, in *pb.GetRollRequest) (*pb.GetRollResponse, error) {
	riconn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Fialed to connect to randint service")
	}

	defer riconn.Close()

	riclient := randint.NewRandintServiceClient(riconn)
	rictx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result := make([]int32, in.R)

	for i := int32(0); i < in.R; i++ {
		r,err := riclient.GetRandint(rictx, &randint.GetRandintRequest{Min: 1, Max: in.D})

		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "Fialed to get roll value from randint service")
		}

		result[i] = r.Value
	}

	return &pb.GetRollResponse{Roll: result}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterDnDServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}