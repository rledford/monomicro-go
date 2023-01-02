package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/rledford/monomicro/randint/proto/v1"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnsafeRandintServiceServer
}

func (s *server) GetRandint(ctx context.Context, in *pb.GetRandintRequest) (*pb.GetRandintResponse, error) {
	log.Printf("Received: %d -> %d", in.Min, in.Max)
	rand.Seed(time.Now().UnixNano())
	value := rand.Int31n(in.Max - in.Min) + in.Min;
	return &pb.GetRandintResponse{Value: value}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRandintServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}