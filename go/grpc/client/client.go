package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"

	rpc "go-grpc-common"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const grpcAddr = "127.0.0.1"
const grpcPort = "9500"
const requests = 10000

func sendRequest(ctx context.Context, client rpc.BenchmarkServiceClient) error {
	response, err := client.Benchmark(ctx, &rpc.Request{Data: "r"})

	if err != nil {
		return err
	}

	response.GetData()

	return nil
}

func runSimpleBenchmark(ctx context.Context, client rpc.BenchmarkServiceClient) error {
	return nil
}

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Server binary:", os.Args[1])

		cmd := exec.Command(os.Args[1])
		cmd.Stdout = os.Stdout

		err := cmd.Start()

		if err != nil {
			panic(err)
		}

		address := fmt.Sprintf("%s:%s", grpcAddr, grpcPort)
		connection, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			panic(err)
		}

		client := rpc.NewBenchmarkServiceClient(connection)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)

		defer cancel()
		defer connection.Close()

		err = sendRequest(ctx, client)
		start := time.Now()

		// Keep retrying until server is up
		for errors.Is(err, syscall.ECONNREFUSED) {
			if time.Now().Sub(start).Seconds() > 1 {
				panic(err)
			}

			err = sendRequest(ctx, client)
		}

		fmt.Println("First request (startup time):", time.Now().Sub(start).Seconds()*1000.0, "ms")

		runSimpleBenchmark(ctx, client)

		syscall.Kill(cmd.Process.Pid, syscall.SIGINT)
		cmd.Wait()
	} else {
		fmt.Println("Expected one argument")
	}
}
