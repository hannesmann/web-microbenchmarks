package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

const httpAddr = "127.0.0.1"
const httpPort = 9000

func main() {
	fmt.Println("Go Echo server started")
	defer fmt.Println("Go Echo server stopped")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "r")
	})

	err := e.Start(fmt.Sprintf("%s:%d", httpAddr, httpPort))

	if err != nil {
		panic(err)
	}
}
