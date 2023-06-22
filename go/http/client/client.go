package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"net/http"
	"syscall"
	"time"
)

const httpAddr = "127.0.0.1"
const httpPort = "9000"
const requests = 10000

func sendRequest(addr string) error {
	// Send a request
	resp, err := http.Get(addr)
	if err != nil {
		return err
	}

	// Read all bytes
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Close connection and start again
	resp.Body.Close()

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

		start := time.Now()
		address := fmt.Sprintf("http://%s:%s", httpAddr, httpPort)

		err = sendRequest(address)

		// Keep retrying until server is up
		for errors.Is(err, syscall.ECONNREFUSED) {
			if time.Now().Sub(start).Seconds() > 1 {
				panic(err)
			}

			err = sendRequest(address)
		}

		fmt.Println("First request (startup time):", time.Now().Sub(start).Seconds() * 1000.0, "ms")
		
		// Send 10000 requests sequentially
		for i := 1; i < requests; i++ {
			if (i + 1) % 1000 == 0 {
				fmt.Printf("Request: %d/%d\n", i + 1, requests)
			}

			err = sendRequest(address)
			if err != nil {
				panic(err)
			}
		}

		elapsed := time.Now().Sub(start)
		seconds := elapsed.Seconds()
		secondsPerRequest := seconds / float64(requests)

		fmt.Println("Average response time:", secondsPerRequest * 1000.0, "ms")

		syscall.Kill(cmd.Process.Pid, syscall.SIGINT)
		cmd.Wait()
	} else {
		fmt.Println("Expected one argument")
	}
}
