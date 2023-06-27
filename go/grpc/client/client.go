package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const grpcAddr = "127.0.0.1"
const grpcPort = "9500"
const requests = 10000

func sendRequest(connection *grpc.ClientConn) error {
	return nil
}

func runSimpleBenchmark(connection *grpc.ClientConn) error {
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

		defer connection.Close()

		err = sendRequest(connection)
		start := time.Now()

		// Keep retrying until server is up
		for errors.Is(err, syscall.ECONNREFUSED) {
			if time.Now().Sub(start).Seconds() > 1 {
				panic(err)
			}

			err = sendRequest(connection)
		}

		fmt.Println("First request (startup time):", time.Now().Sub(start).Seconds()*1000.0, "ms")

		runSimpleBenchmark(connection)

		syscall.Kill(cmd.Process.Pid, syscall.SIGINT)
		cmd.Wait()
	} else {
		fmt.Println("Expected one argument")
	}
}
