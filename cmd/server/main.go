package main

import (
	"context"
	"fmt"
	"net"

	"atomgit.com/cylixlee/protobuf-example-go/generated/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type HelloServiceImpl struct {
	api.UnimplementedHelloServiceServer
}

func (HelloServiceImpl) SayHello(ctx context.Context, r *api.HelloRequest) (*api.HelloResponse, error) {
	fmt.Println("Received:", r)
	return &api.HelloResponse{
		Name:    r.Name,
		Message: "Hello!",
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":11451")
	if err != nil {
		grpclog.Fatalln(err.Error())
	}

	server := grpc.NewServer()
	api.RegisterHelloServiceServer(server, HelloServiceImpl{})

	fmt.Println("gRPC Server running...")

	if err := server.Serve(listen); err != nil {
		grpclog.Fatalln(err.Error())
	}
}
