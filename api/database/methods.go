package database

import (
	"database/sql"
	"fmt"
)

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func GetTasks(db *sql.DB) ([]Task, error) {
	getTasksQuery := "SELECT * from tasks ORDER BY id DESC"

	rows, err := db.Query(getTasksQuery)
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

func CreateTask(db *sql.DB, task Task) error {
	insertTaskQuery := "INSERT INTO tasks (title, description) VALUES (?, ?)"

	_, err := db.Exec(insertTaskQuery, task.Title, task.Description)
	if err != nil {
		return fmt.Errorf("error creating task: %v", err)
	}

	return nil
}

func TaskExists(db *sql.DB, taskId int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM tasks WHERE id = ?)"

	err := db.QueryRow(query, taskId).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func DeleteTask(db *sql.DB, taskId int) error {
	deleteTaskQuery := "DELETE FROM tasks WHERE id = ?"

	_, err := db.Exec(deleteTaskQuery, taskId)
	if err != nil {
		return fmt.Errorf("error creating task: %v", err)
	}

	return nil
}