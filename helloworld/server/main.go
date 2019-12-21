package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/hanjunlee/jira-537/helloworld/helloworld"
)

const (
	port = ":50051"
)

type server struct { 
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("recieve from %s", req.Name)
	return &pb.HelloResponse{
		Message: "hello "+req.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	log.Print("start the server ...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
