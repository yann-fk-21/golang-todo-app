package types

import "time"

type TaskStore interface {
	CreateTask(Task) error
	GetTasks() ([]Task, error)
	// GetTaskById(int) Task
	// UpdateTask(int) error
	// DeleteTask(int) error

}

type Task struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status bool `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}