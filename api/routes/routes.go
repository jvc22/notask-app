package routes

import (
	"log"
	"notask-app/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db database.Database) {
	app.Route("/", func(api fiber.Router) {
		api.Get("/tasks", func(c *fiber.Ctx) error {
			log.Println("GET /tasks")

			tasks, err := db.GetTasks()
			if err != nil {
				log.Printf("> Error retrieving tasks: %v", err)

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error listing tasks.",
				})
			}

			c.Set("Content-Type", "application/json")

			if len(tasks) == 0 {
				return c.JSON(fiber.Map{"tasks": []database.Task{}})
			}
		
			return c.JSON(fiber.Map{"tasks": tasks})
		})

		app.Post("/tasks", func(c *fiber.Ctx) error {
			log.Println("POST /tasks")

			var newTask database.Task
			if err := c.BodyParser(&newTask); err != nil {
				log.Printf("> Error parsing body: %v", err)

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error parsing body.",
				})
			}

			if newTask.Title == "" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Title is required.",
				})
			}

			err := db.CreateTask(newTask)
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

			taskExists, err := db.TaskExists(taskIdAsInt)
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

			err = db.DeleteTask(taskIdAsInt)
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