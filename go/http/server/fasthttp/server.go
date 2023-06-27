package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

const httpAddr = "127.0.0.1"
const httpPort = 9000

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	if string(ctx.Path()) == "/" {
		fmt.Fprintf(ctx, "r")
	} else {
		panic(fmt.Errorf("Invalid path %s", ctx.Path()))
	}
}

func main() {
	fmt.Println("Go fasthttp server started")
	defer fmt.Println("Go fasthttp server stopped")

	err := fasthttp.ListenAndServe(fmt.Sprintf("%s:%d", httpAddr, httpPort), fastHTTPHandler)

	if err != nil {
		panic(err)
	}
}
