package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const httpAddr = "127.0.0.1"
const httpPort = 9000

func main() {
	fmt.Println("Go Gin server started")
	defer fmt.Println("Go Gin server stopped")

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "r")
	})

	err := r.Run(fmt.Sprintf("%s:%d", httpAddr, httpPort))

	if err != nil {
		panic(err)
	}
}
