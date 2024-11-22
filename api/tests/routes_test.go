package tests

import (
	"bytes"
	"io"
	"net/http/httptest"
	"notask-app/database"
	"notask-app/routes"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	_ "github.com/mattn/go-sqlite3"
)

var app *fiber.App
var db *MockDatabase

var testUserId = "testuserid"
var testUsername = "testusername"

func TestMain(m *testing.M) {
	app = fiber.New()

	db = &MockDatabase{}

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("userId", testUserId)

		return c.Next()
	})

	routes.SetupRoutes(app, db)

	m.Run()
}

func TestSignUp(t *testing.T) {
	t.Run("Create account with empty username or password", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/auth/sign-up", bytes.NewReader([]byte(`{"username":"","password":""}`)))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatal("Error sending request:", err)
		}

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		expectedResponse := `{"message":"Username and key are required."}`
		bodyBytes, _ := io.ReadAll(resp.Body)

		assert.Equal(t, expectedResponse, string(bodyBytes))

		db.AssertExpectations(t)
	})

	t.Run("Username already in use", func(t *testing.T) {
		db.On("UserExists", testUsername).Return(true, nil).Once()

		req := httptest.NewRequest("POST", "/auth/sign-up", bytes.NewReader([]byte(`{"username":"` + testUsername + `","password":"testpassword"}`)))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatal("Error sending request:", err)
		}

		assert.Equal(t, fiber.StatusConflict, resp.StatusCode)

		expectedResponse := `{"message":"Username already in use."}`
		bodyBytes, _ := io.ReadAll(resp.Body)

		assert.Equal(t, expectedResponse, string(bodyBytes))

		db.AssertExpectations(t)
	})

	t.Run("Create account successfully", func(t *testing.T) {
		db.On("UserExists", testUsername).Return(false, nil).Once()
		db.On("SignUp", mock.AnythingOfType("database.Auth")).Return(nil).Once()

		req := httptest.NewRequest("POST", "/auth/sign-up", bytes.NewReader([]byte(`{"username":"` + testUsername + `","password":"testpassword"}`)))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatal("Error sending request:", err)
		}

		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

		db.AssertExpectations(t)
	})
}

func TestGetTasks(t *testing.T) {
	t.Run("No tasks in database", func(t *testing.T) {
		tasks := []database.Task{}

		db.On("GetTasks", mock.MatchedBy(func(userId string) bool {
			assert.Equal(t, testUserId, userId)

			return true
		})).Return(tasks, nil).Once()

		req := httptest.NewRequest("GET", "/tasks", nil)

		resp, err := app.Test(req)
		if err != nil {
			t.Fatal("Error sending request:", err)
		}

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		expectedResponse := `{"tasks":[]}`
		bodyBytes, _ := io.ReadAll(resp.Body)

		assert.Equal(t, expectedResponse, string(bodyBytes))

		db.AssertExpectations(t)
	})

	t.Run("Tasks found in database", func(t *testing.T) {
		tasks := []database.Task{
			{Id: 0, Title: "Task 1", Description: "Description 1"},
			{Id: 1, Title: "Task 2", Description: "Description 2"},
		}

		db.On("GetTasks", mock.MatchedBy(func(userId string) bool {
			assert.Equal(t, testUserId, userId)

			return true
		})).Return(tasks, nil).Once()

		req := httptest.NewRequest("GET", "/tasks", nil)

		resp, err := app.Test(req)
		if err != nil {
			t.Fatal("Error sending request:", err)
		}

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		expectedResponse := `{"tasks":[{"id":0,"title":"Task 1","description":"Description 1"},{"id":1,"title":"Task 2","description":"Description 2"}]}`
		bodyBytes, _ := io.ReadAll(resp.Body)

		assert.Equal(t, expectedResponse, string(bodyBytes))

		db.AssertExpectations(t)
	})
}

func TestCreateTask(t *testing.T) {
	t.Run("Create task with empty title", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/tasks", bytes.NewReader([]byte(`{"title":"","description":"Task description"}`)))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatal("Error sending request:", err)
		}

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		expectedResponse := `{"message":"Title is required."}`
		bodyBytes, _ := io.ReadAll(resp.Body)

		assert.Equal(t, expectedResponse, string(bodyBytes))

		db.AssertExpectations(t)
	})

	t.Run("Create task successfully", func(t *testing.T) {
		db.On("CreateTask", mock.MatchedBy(func(userId string) bool {
			assert.Equal(t, testUserId, userId)

			return true
		}), mock.AnythingOfType("database.Task")).Return(nil).Once()

		req := httptest.NewRequest("POST", "/tasks", bytes.NewReader([]byte(`{"title":"Test Task","description":"Task description"}`)))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatal("Error sending request:", err)
		}

		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

		db.AssertExpectations(t)
	})
}

func TestDeleteTask(t *testing.T) {
	t.Run("Invalid Task Id", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/tasks/a", nil)

		resp, err := app.Test(req)
		if err != nil {
			t.Fatal("Error sending request:", err)
		}

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		expectedResponse := `{"message":"Task Id should be a valid number."}`
		bodyBytes, _ := io.ReadAll(resp.Body)

		assert.Equal(t, expectedResponse, string(bodyBytes))

		db.AssertExpectations(t)
	})

	t.Run("Task not found", func(t *testing.T) {
		db.On("TaskExists", mock.MatchedBy(func(userId string) bool {
			assert.Equal(t, testUserId, userId)

			return true
		}), 1).Return(false, nil).Once()

		req := httptest.NewRequest("DELETE", "/tasks/1", nil)

		resp, err := app.Test(req)
		if err != nil {
			t.Fatal("Error sending request:", err)
		}

		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

		expectedResponse := `{"message":"Task not found."}`
		bodyBytes, _ := io.ReadAll(resp.Body)

		assert.Equal(t, expectedResponse, string(bodyBytes))

		db.AssertExpectations(t)
	})

	t.Run("Task deleted successfully", func(t *testing.T) {
		db.On("TaskExists", mock.MatchedBy(func(userId string) bool {
			assert.Equal(t, testUserId, userId)

			return true
		}), 1).Return(true, nil).Once()
		db.On("DeleteTask", mock.MatchedBy(func(userId string) bool {
			assert.Equal(t, testUserId, userId)

			return true
		}), 1).Return(nil).Once()

		req := httptest.NewRequest("DELETE", "/tasks/1", nil)

		resp, err := app.Test(req)
		if err != nil {
			t.Fatal("Error sending request:", err)
		}

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		db.AssertExpectations(t)
	})
}
