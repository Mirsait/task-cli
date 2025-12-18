package commands

import (
	"fmt"
	"time"
)

func Add(tasks []Task, text string) ([]Task, error) {
	if text == "" {
		return tasks, fmt.Errorf("task test cannot be empty")
	}
	nextId := 1
	for _, task := range tasks {
		if task.Id >= nextId {
			nextId = task.Id + 1
		}
	}
	task := Task{
		Id:          nextId,
		Status:      Todo,
		Description: text,
		CreatedAt:   time.Now(),
	}
	return append(tasks, task), nil
}
