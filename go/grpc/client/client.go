package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	rpc "go-grpc-common"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const grpcAddr = "127.0.0.1"
const grpcPort = 9500
const requests = 10000

func sendRequest(ctx context.Context, client rpc.BenchmarkServiceClient) error {
	response, err := client.Benchmark(ctx, &rpc.Request{Data: "r"})

	if err != nil {
		return err
	}

	response.GetData()

	return nil
}

func runSimpleBenchmark(ctx context.Context, client rpc.BenchmarkServiceClient) {
	start := time.Now()

	// Send 10000 requests sequentially
	for i := 0; i < requests; i++ {
		if (i+1)%1000 == 0 {
			fmt.Printf("Request: %d/%d\n", i+1, requests)
		}

		err := sendRequest(ctx, client)
		if err != nil {
			panic(err)
		}
	}

	elapsed := time.Now().Sub(start)
	seconds := elapsed.Seconds()
	secondsPerRequest := seconds / float64(requests)

	fmt.Println("Average response time:", secondsPerRequest*1000.0, "ms")
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

		defer syscall.Kill(cmd.Process.Pid, syscall.SIGINT)

		address := fmt.Sprintf("%s:%d", grpcAddr, grpcPort)

		start := time.Now()
		connection, err := grpc.Dial(address,
			grpc.FailOnNonTempDialError(true),
			grpc.WithReturnConnectionError(),
			grpc.WithBlock(),
			grpc.WithTransportCredentials(insecure.NewCredentials()))

		// Keep retrying until server is up
		for err != nil {
			if !strings.Contains(err.Error(), "connection refused") || time.Now().Sub(start).Seconds() > 1 {
				panic(err)
			}

			connection, err = grpc.Dial(address,
				grpc.FailOnNonTempDialError(true),
				grpc.WithReturnConnectionError(),
				grpc.WithBlock(),
				grpc.WithTransportCredentials(insecure.NewCredentials()))
		}

		fmt.Println("Connection (startup time):", time.Now().Sub(start).Seconds()*1000.0, "ms")

		client := rpc.NewBenchmarkServiceClient(connection)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)

		runSimpleBenchmark(ctx, client)

		cancel()
		connection.Close()
	} else {
		fmt.Println("Expected one argument")
	}
}
