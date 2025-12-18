package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

func Load(filename string) ([]Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.WriteFile(filename, []byte("[]"), 0644); err != nil {
				return nil, fmt.Errorf("create file: %w", err)
			}
			return []Task{}, nil
		}
		return nil, fmt.Errorf("read file: %w", err)
	}

	var tasks []Task
	if err = json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("parse JSON: %w", err)
	}
	return tasks, nil
}
