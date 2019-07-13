//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// Package main implements a greeterServer for Greeter service.
package main

import (
	"log"
	"net"
	"strings"

	pb "github.com/peranikov/grpc-sandbox/grpc/lib"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type greeterServer struct{}

func NewGreeterServer() pb.GreeterServer {
	return &greeterServer{}
}

func (s *greeterServer) SayHelloStream(srv pb.Greeter_SayHelloStreamServer) error {
	var names []string
	for {
		req, err := srv.Recv()
		if err != nil {
			break
		}

		names = append(names, req.Name)
	}

	return srv.SendAndClose(&pb.HelloReply{Message: "Hello " + strings.Join(names, ", ")})
}

func main() {
	lis, err := net.Listen("tcp", port)
	log.Printf("started to listen %v\n", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, NewGreeterServer())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
