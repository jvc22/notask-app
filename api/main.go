package main

import (
	"notask-app/auth"
	"notask-app/database"
	"notask-app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		panic(err)
	}

	connection, err := database.StartDatabase()
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
