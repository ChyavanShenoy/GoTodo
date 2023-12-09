package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"be/routes"
)

// StartServer starts the server
func StartServer(port string) {

	router := gin.Default()

	router.GET("/", routes.HomeRoute(router))

	fmt.Println("Server running at port", port)
	router.Run(port)

}
