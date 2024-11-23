package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var db *SQLDatabase

var testUserId string
var testUsername = "testusername"
var testPassword = "testpassword"

var testTaskId int

func TestMain(m *testing.M) {
	dbFilePath := "./volume/test-database.db"

	connection, err := StartDatabase(dbFilePath)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	db = NewSQLDatabase(connection)

	m.Run()

	os.Remove(dbFilePath)
}

func TestUserMethods(t *testing.T) {
	t.Run("Sign up", func(t *testing.T) {
		newUser := Auth{
			Username: testUsername,
			Password: testPassword,
		}

		userExists, err := db.UserExists(newUser.Username)
		
		assert.Equal(t, nil, err)
		assert.Equal(t, false, userExists)

		err = db.SignUp(newUser)

		assert.Equal(t, nil, err)
	})

	t.Run("Sign in", func(t *testing.T) {
		user := Auth{
			Username: testUsername,
			Password: testPassword,
		}

		userExists, err := db.UserExists(user.Username)

		assert.Equal(t, nil, err)
		assert.Equal(t, true, userExists)

		testUserId, err = db.SignIn(user)

		assert.Equal(t, nil, err)
		assert.NotEmpty(t, testUserId)
	})

	t.Run("Get profile", func(t *testing.T) {
		user, err := db.GetUserProfile(testUserId)

		assert.Equal(t, nil, err)
		assert.Equal(t, testUsername, user.Username)
	})
}

func TestTasksMethods(t *testing.T) {
	t.Run("Create tasks", func(t *testing.T) {
		tasks := []Task{
			{Title: "Task 1", Description: "Description 1"},
			{Title: "Task 2", Description: "Description 2"},
		}

		for _, task := range tasks {
			err := db.CreateTask(testUserId, task)

			assert.Equal(t, nil, err)
		}
	})

	t.Run("Get tasks", func(t *testing.T) {
		tasks, err := db.GetTasks(testUserId)

		assert.Equal(t, nil, err)
		assert.Equal(t, 2, len(tasks))

		testTaskId = tasks[0].Id
	})

	t.Run("Delete task", func(t *testing.T) {
		taskExists, err := db.TaskExists(testUserId, testTaskId)

		assert.Equal(t, nil, err)
		assert.Equal(t, true, taskExists)

		err = db.DeleteTask(testUserId, testTaskId)

		assert.Equal(t, nil, err)

		tasks, err := db.GetTasks(testUserId)

		assert.Equal(t, nil, err)
		assert.Equal(t, 1, len(tasks))
	})
}