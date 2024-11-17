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
	getTasksQuery := "SELECT * from tasks"

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
