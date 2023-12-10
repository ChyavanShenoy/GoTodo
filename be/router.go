package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"be/routes"
)

func StartServer(port string) {
	origin := "*"
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Next()
	})
	router.GET("/", routes.HomeRoute(router))
	router.POST("/api/v1/addTodo", routes.AddTodoRoute(router))

	fmt.Println("Server running at port", port)
	router.Run(port)

}
