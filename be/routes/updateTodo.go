package routes

import (
	"be/models"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func UpdateTodoRoute(router *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Method == "OPTIONS" {
			c.JSON(200, gin.H{
				"message": "Preflight request successful",
			})
			return
		}

		var todo models.Todo
		c.BindJSON(&todo)

		db, err := sql.Open("sqlite3", "database/database.sqlite3")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()

		stmt, err := db.Prepare("UPDATE tasks SET title = ?, description = ?, completed = ?, updated_at = ? WHERE id = ?")
		if err != nil {
			fmt.Println(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(todo.Title, todo.Description, todo.Completed, todo.UpdatedAt, todo.ID)
		if err != nil {
			fmt.Println(err)
		}

		c.JSON(200, gin.H{
			"message": "Todo updated successfully",
		})
	}
}
