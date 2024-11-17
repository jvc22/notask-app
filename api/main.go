package main

import (
	"encoding/json"
	"net/http"
	"notask-app/database"
)

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func main() {
	db, err := database.StartDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /tasks", func(w http.ResponseWriter, r *http.Request) {
		tasks := []Task{
			{
				Id:          1,
				Title:       "Homework",
				Description: "Book & Website",
				Completed:   false,
			},
			{
				Id:          2,
				Title:       "Dishes",
				Description: "",
				Completed:   true,
			},
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(tasks); err != nil {
			http.Error(w, "Unable to encode tasks", http.StatusInternalServerError)

			return
		}
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
