package main

import (
	"notask-app/auth"
	"notask-app/database"
	"notask-app/routes"
	"os"

	_ "notask-app/docs"

	"github.com/gofiber/fiber/v2"
)

// @title notask.app API
// @version 1.0.0
// @description To-do app API to manage tasks
// @host localhost:8080
// @BasePath /
func main() {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		panic("JWT_SECRET_KEY is not set")
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
