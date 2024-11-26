package routes

import (
	"notask-app/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App, db database.Database) {
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/swagger/*", swagger.New(swagger.Config{
		DeepLinking:  false,
		DocExpansion: "list",
	}))

	app.Route("/", func(api fiber.Router) {
		handlerGetTasks(api, db)
		handlerCreateTask(api, db)
		handlerDeleteTask(api, db)
	})

	app.Route("/auth", func(api fiber.Router) {
		handlerSignUp(api, db)
		handlerSignIn(api, db)
	})

	app.Route("/user", func(api fiber.Router) {
		handlerGetProfile(api, db)
	})
}
