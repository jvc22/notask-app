package routes

import (
	"database/sql"
	"log"
	"notask-app/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	app.Route("/", func(api fiber.Router) {
		api.Get("/tasks", func(c *fiber.Ctx) error {
			log.Println("GET /tasks")

			tasks, err := database.GetTasks(db)
			if err != nil {
				log.Printf("> Error retrieving tasks: %v", err)

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error listing tasks.",
				})
			}

			c.Set("Content-Type", "application/json")

			return c.JSON(fiber.Map{"tasks": tasks})
		})

		app.Post("/tasks", func(c *fiber.Ctx) error {
			log.Println("POST /tasks")

			var newTask database.Task
			if err := c.BodyParser(&newTask); err != nil {
				if newTask.Title == "" {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"message": "Title is required.",
					})
				}

				log.Printf("> Error parsing body: %v", err)

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error parsing body.",
				})
			}

			err := database.CreateTask(db, newTask)
			if err != nil {
				log.Printf("> Error creating task: %v", err)

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error creating task.",
				})
			}

			return c.SendStatus(fiber.StatusCreated)
		})

		api.Delete("/tasks/:id", func(c *fiber.Ctx) error {
			log.Println("DELETE /tasks")

			taskId := c.Params("id")
			if taskId == "" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Task Id is required.",
				})
			}

			taskIdAsInt, err := strconv.Atoi(taskId)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Task Id should be a valid number.",
				})
			}

			taskExists, err := database.TaskExists(db, taskIdAsInt)
			if err != nil {
				log.Printf("> Error checking task existence: %v", err)

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error checking task existence.",
				})
			}

			if !taskExists {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"message": "Task not found.",
				})
			}

			err = database.DeleteTask(db, taskIdAsInt)
			if err != nil {
				log.Printf("> Error deleting task: %v", err)

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error deleting task.",
				})
			}
			
			return c.SendStatus(fiber.StatusOK)
		})
	})
}