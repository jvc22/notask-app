package main

import (
	"net/http"
	"notask-app/database"
	"notask-app/routes"
)

func main() {
	db, err := database.StartDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	routes.SetupRoutes(mux, db)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
