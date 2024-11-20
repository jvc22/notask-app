package database

import "github.com/stretchr/testify/mock"

type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) GetTasks() ([]Task, error) {
	args := m.Called()

	return args.Get(0).([]Task), args.Error(1)
}

func (m *MockDatabase) CreateTask(task Task) error {
	args := m.Called(task)

	return args.Error(0)
}

func (m *MockDatabase) DeleteTask(id int) error {
	args := m.Called(id)

	return args.Error(0)
}

func (m *MockDatabase) TaskExists(id int) (bool, error) {
	args := m.Called(id)
	
	return args.Bool(0), args.Error(1)
}
