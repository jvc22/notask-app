package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var connection *sql.DB

func NewSQLDatabase(connection *sql.DB) *SQLDatabase {
	return &SQLDatabase{connection}
}

func createUsersTable() error {
	createUsersTableSQL := `CREATE TABLE IF NOT EXISTS users (
		"id" TEXT PRIMARY KEY,
		"username" TEXT,
		"password" TEXT
	);`

	_, err := connection.Exec(createUsersTableSQL)
	if err != nil {
		return err
	}

	return nil
}

func createTasksTable() error {
	createTasksTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"description" TEXT,
		"userId" TEXT,
		FOREIGN KEY (userId) REFERENCES users(id)
	);`

	_, err := connection.Exec(createTasksTableSQL)
	if err != nil {
		return err
	}

	return nil
}

func StartDatabase(path string) (*sql.DB, error) {
	var err error

	connection, err = sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	err = createUsersTable()
	if err != nil {
		return nil, err
	}

	err = createTasksTable()
	if err != nil {
		return nil, err
	}

	return connection, nil
}
