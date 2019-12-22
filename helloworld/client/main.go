package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/hanjunlee/jira-537/helloworld/helloworld"
	"google.golang.org/grpc"
)

const (
	defaultName = "world"
)

var (
	waitSecond = flag.Uint("wait", 0, "The time of waiting the next call. If the value is 0 the client request only one time.")
	address    = flag.String("addr", "localhost", "The address of the server.")
	port       = flag.String("port", "9000", "The port of the server.")
)

type client struct {
	gc pb.GreeterClient
}

func (c *client) SayHello(ws uint) {
	if ws <= 0 {
		r, err := c.sayHello()
		if err != nil {
			log.Printf("error from the server: %v", err)
		}

		log.Printf("greeting: %s", r)
		return
	}

	ticker := time.NewTicker(time.Duration(ws) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			log.Printf("current time: %s", t)
			r, err := c.sayHello()
			if err != nil {
				log.Printf("error from the server: %v", err)
			}

			log.Printf("greeting: %s", r)
		}
	}
}

func (c *client) sayHello() (*pb.HelloResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return c.gc.SayHello(ctx, &pb.HelloRequest{Name: defaultName})
}

func main() {
	flag.Parse()

	// Set up a connection to the server.
	log.Printf("open the connection to %s", *address+":"+*port)
	conn, err := grpc.Dial(*address+":"+*port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := &client{
		gc: pb.NewGreeterClient(conn),
	}
	c.SayHello(*waitSecond)
}
