package main

import (
	"notask-app/database"
)

func main() {
	db := database.StartDatabase()

	defer db.Close()
}
