package main

import (
	"log"
	"notask-app/database"
)

func main() {
	db, err := database.StartDatabase()
	if err != nil {
		log.Fatal("Error during database start:", err)
	}

	defer db.Close()
}
