package database

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Database interface {
	UserExists(username string) (bool, error)
	SignUp(data Auth) error
	SignIn(data Auth) (string, error)
	GetTasks(userId string) ([]Task, error)
	TaskExists(userId string, id int) (bool, error)
	CreateTask(userId string, task Task) error
	DeleteTask(userId string, id int) error
}
