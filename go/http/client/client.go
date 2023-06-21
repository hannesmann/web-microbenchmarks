package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Server binary:", os.Args[1])

		cmd := exec.Command(os.Args[1])
		err := cmd.Start()

		if err != nil {
			panic(err)
		}

		cmd.Wait()
	} else {
		fmt.Println("Expected one argument")
	}
}
