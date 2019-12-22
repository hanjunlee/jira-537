package main

import (
	"context"
	"flag"
	"log"
	"net"
	"time"

	pb "github.com/hanjunlee/jira-537/helloworld/helloworld"
	"google.golang.org/grpc"
)

var (
	port = flag.String("port", "9000", "the port of the server.")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("recieve from %s", req.Name)
	time.Sleep(2*time.Second)
	return &pb.HelloResponse{
		Message: "hello " + req.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("start the server on %s port ...", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
