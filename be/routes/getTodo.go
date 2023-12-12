package routes

import (
	"be/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func GetTodosRoute(router *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("sqlite3", "database/database.sqlite3")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		rows, err := db.Query("SELECT * FROM tasks")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var todos []models.Todo
		for rows.Next() {
			var todo models.Todo
			err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
			if err != nil {
				log.Fatal(err)
			}
			todos = append(todos, todo)
		}

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(todos)

		c.JSON(http.StatusOK, todos)
	}
}
