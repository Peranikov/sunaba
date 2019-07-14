//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// Package main implements a greeterServer for Greeter service.
package grpc

import (
	"io"
	"log"
	"net"
	"strings"

	pb "github.com/peranikov/sunaba/grpc/lib"
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
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		names = append(names, req.Name)
	}

	return srv.SendAndClose(&pb.HelloReply{Message: "Hello " + strings.Join(names, ", ")})
}

func Run() error {
	lis, err := net.Listen("tcp", port)
	log.Printf("started gRPC to listen %v\n", port)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, NewGreeterServer())
	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
