package main

import (
	"context"
	"github.com/tkxkd0159/buf-proto/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	addr := "localhost:3000"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Can not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "LJS"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("From gRPC srv : %s", r.GetMessage())

	r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: "New Name"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("From gRPC srv : %s", r.GetMessage())
}
