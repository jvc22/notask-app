package main

import (
	"notask-app/database"
	"notask-app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db, err := database.StartDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET, POST, DELETE",
	}))

	routes.SetupRoutes(app, db)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
