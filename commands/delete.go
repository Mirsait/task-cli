package commands

import (
	"fmt"
)

func Delete(tasks []Task, id int) ([]Task, error) {
	for j, task := range tasks {
		if task.Id == id {
			return append(tasks[:j], tasks[j+1:]...), nil
		}
	}
	return tasks, fmt.Errorf("The task with ID:%d not found.\n", id)
}
