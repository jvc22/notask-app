package routes

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"notask-app/database"
	"strconv"
)

func SetupRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("GET /tasks", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET /tasks")

		tasks, err := database.GetTasks(db)
		if err != nil {
			http.Error(w, "Error listing tasks", http.StatusInternalServerError)

			return
		}

		w.Header().Set("Content-Type", "application/json")

		response := map[string]interface{}{
			"tasks": tasks,
		}

		if err = json.NewEncoder(w).Encode(response); err != nil {
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

		err := database.CreateTask(db, newTask)
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

		exists, err := database.TaskExists(db, taskIdAsInt)
		if err != nil {
			http.Error(w, "Error checking task existence", http.StatusInternalServerError)

			return
		}

		if !exists {
			http.Error(w, "Task not found", http.StatusNotFound)

			return
		}

		err = database.DeleteTask(db, taskIdAsInt)
		if err != nil {
			http.Error(w, "Error deleting task", http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
