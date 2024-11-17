package main

import (
	"net/http"
	"notask-app/database"
	"notask-app/routes"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)

			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	db, err := database.StartDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	routes.SetupRoutes(mux, db)

	handlerWithCORS := corsMiddleware(mux)

	if err := http.ListenAndServe(":8080", handlerWithCORS); err != nil {
		panic(err)
	}
}
