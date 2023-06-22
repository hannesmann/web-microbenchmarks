package main

import (
	"fmt"
	"net/http"
)

const httpAddr = "127.0.0.1"
const httpPort = "9000"

func main() {
	fmt.Println("Go net/http server started")
	defer fmt.Println("Go net/http server stopped")

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprint(writer, "r")
	})

	http.ListenAndServe(fmt.Sprintf("%s:%s", httpAddr, httpPort), nil)
}