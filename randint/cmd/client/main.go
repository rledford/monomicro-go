package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/rledford/monomicro/randint/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultMin = 0;
	defaultMax = 100;
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	min = flag.Int("min", defaultMin, "Min range")
	max = flag.Int("max", defaultMax, "Max range")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRandintServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetRandint(ctx, &pb.GetRandintRequest{Min: int32(*min), Max: int32(*max)})
	if err != nil {
		log.Fatalf("could not get value: %v", err)
	}
	log.Printf("Value: %d", r.GetValue())
}