package functions

import (
	"be/models"
	"database/sql"
	"log"
)

func AddTodoToDB(data models.Todo) error {
	db, err := sql.Open("sqlite3", "database/database.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Prepare the SQL statement
	stmt, err := db.Prepare("INSERT INTO todos(title, description) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement with the passed data
	_, err = stmt.Exec(data.Title, data.Description)
	if err != nil {
		return err
	}

	return nil
}
