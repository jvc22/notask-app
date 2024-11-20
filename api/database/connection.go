package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var connection *sql.DB

func NewSQLDatabase(connection *sql.DB) *SQLDatabase {
	return &SQLDatabase{connection}
}

func createTasksTable() error {
	createTasksTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"description" TEXT
	);`

	_, err := connection.Exec(createTasksTableSQL)
	if err != nil {
		return err
	}

	return nil
}

func StartDatabase() (*sql.DB, error) {
	var err error

	connection, err = sql.Open("sqlite3", "./database/volume/tasks.db")
	if err != nil {
		return nil, err
	}

	err = createTasksTable()
	if err != nil {
		return nil, err
	}

	return connection, nil
}