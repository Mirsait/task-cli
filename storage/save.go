package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

func Save(filename string, tasks []Task) error {
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("encode tasks: %w", err)
	}
	return os.WriteFile(filename, jsonData, 0644)
}
