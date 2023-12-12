package models

import (
	"database/sql"
)

type Todo struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Completed   bool         `json:"completed"`
	CreatedAt   sql.NullTime `json:"createdAt"`
	UpdatedAt   sql.NullTime `json:"updatedAt"`
}
