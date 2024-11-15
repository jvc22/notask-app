package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func createTasksTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"description" TEXT,
		"completed" BOOLEAN
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

func StartDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./database/tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	createTasksTable(db)

	fmt.Println("SQLite3 database started 🔥")

	return db
}
