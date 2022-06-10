package grpcreg

import (
	"context"
	"fmt"
	"github.com/tkxkd0159/buf-proto/grpc/pb"
	"log"
)

func showStartMsg() {
	fmt.Println("Grpc server is running...")
}

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello again, " + in.GetName()}, nil
}

func helloSrv() *server {
	s := &server{}
	return s
}
