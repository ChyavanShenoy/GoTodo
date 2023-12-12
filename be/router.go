package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"be/routes"
)

func StartServer(port string) {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Next()

		if c.Request.Method == "OPTIONS" {
			c.JSON(200, gin.H{
				"message": "Preflight request successful",
			})
			return
		}
	})
	router.GET("/", routes.HomeRoute(router))

	router.GET("/api/v1/getTodos", routes.GetTodosRoute(router))
	router.POST("/api/v1/addTodo", routes.AddTodoRoute(router))
	router.PUT("/api/v1/updateTodo", routes.UpdateTodoRoute(router))
	// router.DELETE("/api/v1/deleteTodo", routes.DeleteTodoRoute(router))

	fmt.Println("Server running at port", port)
	router.Run(port)

}
