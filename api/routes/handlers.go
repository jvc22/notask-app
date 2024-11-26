package routes

import (
	"log"
	"notask-app/auth"
	"notask-app/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Summary Sign in user
// @Description Authenticate user based on credentials
// @Tags auth
// @Param auth body database.Auth true "Auth object"
// @Accept json
// @Produce json
// @Success 200 "Accessed"
// @Failure 400 {object} ResponseErrorMessage "Username and password are required"
// @Failure 404 {object} ResponseErrorMessage "Username not found or invalid password"
// @Failure 500 {object} ResponseErrorMessage "Error parsing body, searching user, or generating token"
// @Router /auth/sign-in [post]
func handlerSignIn(api fiber.Router, db database.Database) {
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
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"message": "Invalid password.",
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

		c.Locals("userId", userId)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token": token,
		})
	})
}

// @Summary Sign up user
// @Description Create new user account on database
// @Tags auth
// @Param auth body database.Auth true "Auth object"
// @Accept json
// @Produce json
// @Success 201 "Created"
// @Failure 400 {object} ResponseErrorMessage "Username and password are required"
// @Failure 409 {object} ResponseErrorMessage "Username already in use"
// @Failure 500 {object} ResponseErrorMessage "Error parsing body or registering new account"
// @Router /auth/sign-up [post]
func handlerSignUp(api fiber.Router, db database.Database) {
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
				"message": "Username and password are required.",
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
}

// @Summary Fetch and return user info
// @Description Get all profile user information registered on the database
// @Tags user
// @Produce json
// @Success 200 "Fetched"
// @Failure 500 {object} ResponseErrorMessage "User not found"
// @Failure 500 {object} ResponseErrorMessage "Error fetching user profile"
// @Router /user [get]
func handlerGetProfile(api fiber.Router, db database.Database) {
	api.Get("/profile", func(c *fiber.Ctx) error {
		log.Println("GET /user/profile")

		userId := c.Locals("userId").(string)

		user, err := db.GetUserProfile(userId)
		if err != nil {
			log.Printf("> Authentication error: %v", err)

			if err == database.ErrUserNotFound {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"message": "User not found.",
				})
			}

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error. Try again in a few minutes.",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"username": user.Username,
		})
	})
}

// @Summary Fetch and return tasks
// @Description Get all tasks created specifically by the user
// @Tags tasks
// @Produce json
// @Success 200 {object} []database.Task "Fetched"
// @Failure 500 {object} ResponseErrorMessage "Error parsing body or fetching tasks"
// @Router /tasks [get]
func handlerGetTasks(api fiber.Router, db database.Database) {
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
}

// @Summary Create new task
// @Description Insert new user task into database
// @Tags tasks
// @Param task body database.Task true "Task object"
// @Accept json
// @Produce json
// @Success 201 "Created"
// @Failure 400 {object} ResponseErrorMessage "Title is required"
// @Failure 500 {object} ResponseErrorMessage "Error parsing body or creating task"
// @Router /tasks [post]
func handlerCreateTask(api fiber.Router, db database.Database) {
	api.Post("/tasks", func(c *fiber.Ctx) error {
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
}

// @Summary Delete task
// @Description Remove task from database using id
// @Tags tasks
// @Param id path string true "Task Id"
// @Produce json
// @Success 200 "Deleted"
// @Failure 400 {object} ResponseErrorMessage "Task Id is required"
// @Failure 400 {object} ResponseErrorMessage "Task Id should be a valid number | Task not found"
// @Failure 500 {object} ResponseErrorMessage "Error checking or deleting task"
// @Router /tasks [delete]
func handlerDeleteTask(api fiber.Router, db database.Database) {
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
}
