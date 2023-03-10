package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/rledford/monomicro-go/randint/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnsafeRandintServiceServer
}

func (s *server) GetRandint(ctx context.Context, in *pb.GetRandintRequest) (*pb.GetRandintResponse, error) {
	if (in.Min < 0 || in.Min >= in.Max) {
		return nil, status.Error(codes.InvalidArgument, "invalid min or max")
	}

	log.Printf("Received: %d -> %d", in.Min, in.Max)

	rand.Seed(time.Now().UnixNano())
	value := rand.Int31n(in.Max - in.Min + 1) + in.Min;

	return &pb.GetRandintResponse{Value: value}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterRandintServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}