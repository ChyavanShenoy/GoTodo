package functions

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	tableName       = "tasks"
	databasePath    = "database/database.sqlite3"
	createTableStmt = `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		title TEXT, 
		description TEXT, 
		completed BOOLEAN, 
		created_at DATETIME, 
		updated_at DATETIME
	)`
)

func CheckCreateDB() error {
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	defer db.Close()

	if err := createTable(db); err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}

	return nil
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(createTableStmt)
	if err != nil {
		return fmt.Errorf("error executing create table statement: %v", err)
	}

	log.Printf("Table %s checked/created successfully\n", tableName)
	return nil
}
