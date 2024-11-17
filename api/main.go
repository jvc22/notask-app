package main

import (
	"encoding/json"
	"log"
	"net/http"
	"notask-app/database"
	"strconv"
)

func main() {
	db, err := database.StartDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /tasks", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET /tasks")

		tasks, err := database.GetTasks(db)
		if err != nil {
			http.Error(w, "Error listing tasks", http.StatusInternalServerError)

			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err = json.NewEncoder(w).Encode(tasks); err != nil {
			http.Error(w, "Unable to encode tasks", http.StatusInternalServerError)

			return
		}
	})

	mux.HandleFunc("POST /tasks", func(w http.ResponseWriter, r *http.Request) {
		log.Println("POST /tasks")

		var newTask database.Task
		if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)

			return
		}

		err = database.CreateTask(db, newTask)
		if err != nil {
			http.Error(w, "Error creating task", http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	mux.HandleFunc("DELETE /tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		log.Println("DELETE /tasks")

		taskId := r.PathValue("id")
		if taskId == "" {
			http.Error(w, "Task ID is required", http.StatusBadRequest)

			return
		}

		taskIdAsInt, err := strconv.Atoi(taskId)
		if err != nil {
			http.Error(w, "Invalid Task ID", http.StatusBadRequest)

			return
		}

		err = database.DeleteTask(db, taskIdAsInt)
		if err != nil {
			http.Error(w, "Error deleting task", http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusOK)
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
