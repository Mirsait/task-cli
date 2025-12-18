package commands

import (
	"fmt"
	"time"
)

func Mark(tasks []Task, id int, status Status) ([]Task, error) {
	for j, v := range tasks {
		if v.Id == id {
			tasks[j].Status = status
			tasks[j].UpdatedAt = time.Now()
			return tasks, nil
		}
	}
	return tasks, fmt.Errorf("The task with ID:%d not found.\n", id)
}
