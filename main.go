package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server")

	server := gin.Default()
	GetRoutes(server)

	server.Run(":8080")
}
