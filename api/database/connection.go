package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func createTasksTable(db *sql.DB) error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"description" TEXT,
		"completed" BOOLEAN
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	return nil
}

func StartDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database/tasks.db")
	if err != nil {
		return nil, err
	}

	err = createTasksTable(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
