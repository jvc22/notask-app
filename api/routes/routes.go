package routes

import (
	"log"
	"notask-app/auth"
	"notask-app/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db database.Database) {
	app.Route("/", func(api fiber.Router) {
		api.Get("/tasks", func(c *fiber.Ctx) error {
			log.Println("GET /tasks")

			userId := c.Locals("userId").(string)

			tasks, err := db.GetTasks(userId)
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

			userId := c.Locals("userId").(string)

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

			err := db.CreateTask(userId, newTask)
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

			userId := c.Locals("userId").(string)

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

			taskExists, err := db.TaskExists(userId, taskIdAsInt)
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

			err = db.DeleteTask(userId, taskIdAsInt)
			if err != nil {
				log.Printf("> Error deleting task: %v", err)

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error deleting task.",
				})
			}

			return c.SendStatus(fiber.StatusOK)
		})
	})

	app.Route("/auth", func(api fiber.Router) {
		api.Post("/sign-up", func(c *fiber.Ctx) error {
			log.Println("POST /auth/sign-up")

			var signUpData database.Auth
			if err := c.BodyParser(&signUpData); err != nil {
				log.Printf("> Error parsing body: %v", err)

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error. Try again in a few minutes",
				})
			}

			if signUpData.Username == "" || signUpData.Password == "" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Username and key are required.",
				})
			}

			userExists, err := db.UserExists(signUpData.Username)
			if err != nil {
				log.Printf("> Error searching user: %v", err)

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error sarching user.",
				})
			}

			if userExists {
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{
					"message": "Username already in use.",
				})
			}

			err = db.SignUp(signUpData)
			if err != nil {
				log.Printf("> Error registering new account: %v", err)

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error. Try again in a few minutes.",
				})
			}

			return c.SendStatus(fiber.StatusCreated)
		})

		api.Post("/sign-in", func(c *fiber.Ctx) error {
			log.Println("POST /auth/sign-in")

			var signInData database.Auth
			if err := c.BodyParser(&signInData); err != nil {
				log.Printf("> Error parsing body: %v", err)

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error. Try again in a few minutes",
				})
			}

			if signInData.Username == "" || signInData.Password == "" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Username and password are required.",
				})
			}

			userExists, err := db.UserExists(signInData.Username)
			if err != nil {
				log.Printf("> Error searching user: %v", err)

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error. Try again in a few minutees.",
				})
			}

			if !userExists {
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{
					"message": "Username not found.",
				})
			}

			userId, err := db.SignIn(signInData)
			if err != nil {
				log.Printf("> Authentication error: %v", err)

				if err == database.ErrInvalidCredentials {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"message": "Invalid key.",
					})
				}

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error. Try again in a few minutes.",
				})
			}

			token, err := auth.GenerateJWT(userId)
			if err != nil {
				log.Printf("> Error generating JWT: %v", err)

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error. Try again in a few minutes.",
				})
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"token": token,
			})
		})
	})
}
