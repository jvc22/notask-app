package db

import (
	"database/sql"
	"fmt"
	"log"
)

func StartDatabase() {
	db, err := sql.Open("sqlite3", "./db/tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"description" TEXT,
		"completed" BOOLEAN
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("error during the creation of table:", err)
	}

	fmt.Println("SQLite3 database started 🔥")

	defer db.Close()
}
