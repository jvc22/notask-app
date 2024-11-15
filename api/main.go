package main

import (
	"log"
	"notask-app/database"
)

func main() {
	db, err := database.StartDatabase()
	if err != nil {
		log.Fatal("Error during opening database:", err)
	}

	defer db.Close()
}
