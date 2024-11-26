package main

import (
	"notask-app/auth"
	"notask-app/database"
	"notask-app/routes"

	_ "notask-app/docs"

	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
)

// @title notask.app API
// @version 1.0.0
// @description To-do app API to manage tasks
// @host localhost:8080
// @BasePath /
func main() {
	if err := godotenv.Load("/app/.env"); err != nil {
		if err := godotenv.Load("../.env"); err != nil {
			panic(err)
		}
	}

	connection, err := database.StartDatabase("./database/volume/database.db")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	db := database.NewSQLDatabase(connection)
	app := fiber.New()

	auth.SetCORS(app)
	auth.AuthMiddleware(app, connection)
	routes.SetupRoutes(app, db)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
