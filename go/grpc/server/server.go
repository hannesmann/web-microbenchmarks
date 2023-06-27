package main

import (
	"context"
	"fmt"
	"net"

	rpc "go-grpc-common"

	"google.golang.org/grpc"
)

const grpcAddr = "127.0.0.1"
const grpcPort = 9500

type Server struct {
	rpc.UnimplementedBenchmarkServiceServer
}

func (s *Server) Benchmark(ctx context.Context, in *rpc.Request) (*rpc.Response, error) {
	return &rpc.Response{Data: "r"}, nil
}

func main() {
	fmt.Println("Go gRPC server started")
	defer fmt.Println("Go gRPC server stopped")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", grpcAddr, grpcPort))

	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	rpc.RegisterBenchmarkServiceServer(server, &Server{})

	err = server.Serve(listener)

	if err != nil {
		panic(err)
	}
}
