package grpcreg

import (
	"fmt"
	"github.com/tkxkd0159/buf-proto/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartGrpcSrv() {
	gPort := 3000
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", gPort))
	if err != nil {
		log.Fatalf("Fail to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, helloSrv())
	showStartMsg()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
