package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"net/http"
	"time"
)

const httpAddr = "127.0.0.1"
const httpPort = "9000"

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Server binary:", os.Args[1])

		cmd := exec.Command(os.Args[1])
		err := cmd.Start()

		if err != nil {
			panic(err)
		}

		// Wait for server to get ready
		time.Sleep(3 * time.Second)

		// Send 10000 requests sequentially
		start := time.Now()
		address := fmt.Sprintf("http://%s:%s", httpAddr, httpPort)

		for i := 0; i < 10000; i++ {
			// Send a request
			resp, err := http.Get(address)
			if err != nil {
				panic(err)
			}
			// Read all bytes
			_, err = io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			// Close connection and start again
			resp.Body.Close()
		}

		elapsed := time.Now().Sub(start)
		seconds := elapsed.Seconds()
		secondsPerRequest := seconds / 10000.0

		fmt.Println("Average response time:", secondsPerRequest * 1000.0, "ms")

		cmd.Wait()
	} else {
		fmt.Println("Expected one argument")
	}
}
