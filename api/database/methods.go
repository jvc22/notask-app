package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type SQLDatabase struct {
	*sql.DB
}

var ErrInvalidCredentials = fmt.Errorf("invalid credentials")
var ErrUserNotFound = fmt.Errorf("user not found")

func (db *SQLDatabase) UserExists(username string) (bool, error) {
	var exists bool

	userExistsQuery := "SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)"

	err := db.QueryRow(userExistsQuery, username).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (db *SQLDatabase) SignUp(data Auth) error {
	newUserId := uuid.New()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error encrypting key: %v", err)
	}

	createAccountQuery := "INSERT INTO users (id, username, password) VALUES (?, ?, ?)"

	_, err = db.Exec(createAccountQuery, newUserId, data.Username, hashPassword)
	if err != nil {
		return fmt.Errorf("error creating password: %v", err)
	}

	return nil
}

func (db *SQLDatabase) SignIn(data Auth) (string, error) {
	getUserDataQuery := "SELECT id, password FROM users WHERE username = ?"

	var userId, hashPassword string

	err := db.QueryRow(getUserDataQuery, data.Username).Scan(&userId, &hashPassword)
	if err != nil {
		return "", fmt.Errorf("error retrieving user data: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(data.Password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	return userId, nil
}

func (db *SQLDatabase) GetUserProfile(userId string) (User, error) {
	getUserProfileQuery := "SELECT username FROM users WHERE id = ?"

	var user User

	err := db.QueryRow(getUserProfileQuery, userId).Scan(&user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, ErrUserNotFound
		}

		return User{}, fmt.Errorf("error fetching user profile: %v", err)
	}

	return user, nil
}

func (db *SQLDatabase) GetTasks(userId string) ([]Task, error) {
	getTasksQuery := "SELECT id, title, description FROM tasks WHERE userId = ? ORDER BY id DESC"

	rows, err := db.Query(getTasksQuery, userId)
	if err != nil {
		return nil, fmt.Errorf("error fetching tasks: %v", err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.Id, &task.Title, &task.Description); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (db *SQLDatabase) CreateTask(userId string, task Task) error {
	insertTaskQuery := "INSERT INTO tasks (title, description, userId) VALUES (?, ?, ?)"

	_, err := db.Exec(insertTaskQuery, task.Title, task.Description, userId)
	if err != nil {
		return fmt.Errorf("error creating task: %v", err)
	}

	return nil
}

func (db *SQLDatabase) TaskExists(userId string, taskId int) (bool, error) {
	var exists bool

	taskExistsQuery := "SELECT EXISTS(SELECT 1 FROM tasks WHERE id = ? AND userId = ?)"

	err := db.QueryRow(taskExistsQuery, taskId, userId).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (db *SQLDatabase) DeleteTask(userId string, taskId int) error {
	deleteTaskQuery := "DELETE FROM tasks WHERE id = ? AND userId = ?"

	_, err := db.Exec(deleteTaskQuery, taskId, userId)
	if err != nil {
		return fmt.Errorf("error creating task: %v", err)
	}

	return nil
}
