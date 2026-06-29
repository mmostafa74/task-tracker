package main

import (
	"fmt"
	"time"
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func colorStatus(status Status) string {
	switch status {
	case StatusTodo:
		return "\033[31m" + string(status) + "\033[0m" // Red
	case StatusInProgress:
		return "\033[33m" + string(status) + "\033[0m" // Yellow
	case StatusDone:
		return "\033[32m" + string(status) + "\033[0m" // Green
	default:
		return string(status)
	}
}

func (t Task) String() string {
	return fmt.Sprintf("ID: %d, Description: %s, Status: %s, Created At: %s, Updated At: %s", t.ID, t.Description, colorStatus(t.Status), t.CreatedAt.Format("2006-01-02 15:04:05"), t.UpdatedAt.Format("2006-01-02 15:04:05"))
}
