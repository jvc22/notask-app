package main

import (
	"notask-app/db"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db.StartDatabase()
}
