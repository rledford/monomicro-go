package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/rledford/monomicro/dnd/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

type server struct {
	pb.UnsafeDnDServiceServer
}

func (s *server) GetRoll(ctx context.Context, in *pb.GetRollRequest) (*pb.GetRollResponse, error) {
	return &pb.GetRollResponse{Roll: []int32{0}}, nil
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