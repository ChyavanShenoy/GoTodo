package routes

import (
	"be/models"
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddTodoRoute(router *gin.Engine) gin.HandlerFunc {
	fmt.Println("AddTodo route")

	return func(c *gin.Context) {
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			return
		}
		bodyString := string(bodyBytes)
		fmt.Println("Raw request body:", bodyString)

		// Now reassign the body so it can be read again during binding
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		var todo models.Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			fmt.Println("Error binding JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("Todo:", todo)

		c.JSON(http.StatusOK, gin.H{"message": "Todo added"})
	}
}
