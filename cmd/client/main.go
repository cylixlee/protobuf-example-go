package main

import (
	"context"
	"fmt"

	"atomgit.com/cylixlee/protobuf-example-go/generated/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

func main() {
	client, err := grpc.NewClient(":11451", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		grpclog.Fatalln(err.Error())
	}
	defer client.Close()

	serviceClient := api.NewHelloServiceClient(client)
	response, err := serviceClient.SayHello(context.Background(), &api.HelloRequest{
		Name: "Cylix",
	})
	if err != nil {
		grpclog.Fatalln(err.Error())
	}

	fmt.Println("gRPC Server:", response)
}
