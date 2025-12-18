package models

import (
	"time"
)

type Status string

const (
	Todo     Status = "todo"
	Progress Status = "in-progress"
	Done     Status = "done"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func newTask(id int, text string) Task {
	return Task{
		Id:          id,
		Status:      Todo,
		Description: text,
		CreatedAt:   time.Now(),
	}
}
