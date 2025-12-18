package commands

import (
	"fmt"
	"time"
)

func Update(tasks []Task, id int, newText string) ([]Task, error) {
	if newText == "" {
		return tasks, fmt.Errorf("Task text cannot be empty")
	}
	for j, v := range tasks {
		if v.Id == id {
			tasks[j].Description = newText
			tasks[j].UpdatedAt = time.Now()
			return tasks, nil
		}
	}
	return tasks, fmt.Errorf("The task with ID:%d has not found.\n", id)
}
