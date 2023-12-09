package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	PORT := ":8080"

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	fmt.Println("Server running at port", PORT)
	router.Run(PORT)

}
