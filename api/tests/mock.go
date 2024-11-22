package tests

import (
	"notask-app/database"

	"github.com/stretchr/testify/mock"
)

type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) UserExists(username string) (bool, error) {
	args := m.Called(username)

	return args.Bool(0), args.Error(1)
}

func (m *MockDatabase) SignUp(data database.Auth) error {
	args := m.Called(data)

	return args.Error(0)
}

func (m *MockDatabase) SignIn(data database.Auth) (string, error) {
	args := m.Called(data)

	return args.String(0), args.Error(1)
}

func (m *MockDatabase) GetUserProfile(userId string) (database.User, error) {
	args := m.Called(userId)

	return args.Get(0).(database.User), args.Error(1)
}

func (m *MockDatabase) GetTasks(userId string) ([]database.Task, error) {
	args := m.Called(userId)

	return args.Get(0).([]database.Task), args.Error(1)
}

func (m *MockDatabase) CreateTask(userId string, task database.Task) error {
	args := m.Called(userId, task)

	return args.Error(0)
}

func (m *MockDatabase) TaskExists(userId string, id int) (bool, error) {
	args := m.Called(userId, id)

	return args.Bool(0), args.Error(1)
}

func (m *MockDatabase) DeleteTask(userId string, id int) error {
	args := m.Called(userId, id)

	return args.Error(0)
}
