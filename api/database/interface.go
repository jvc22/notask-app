package database

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Database interface {
	GetTasks() ([]Task, error)
	CreateTask(task Task) error
	DeleteTask(id int) error
	TaskExists(id int) (bool, error)
}
