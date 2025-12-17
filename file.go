package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func loadData(filename string) []Task {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := os.WriteFile(filename, []byte("[]"), 0644); err != nil {
			panic(err)
		}
	}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v\n", err)
	}
	return tasks
}

func saveData(filename string, data []Task) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("Error marshaling to JSON: %s\n", err)
		return
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		log.Printf("Error writing to file: %s\n", err)
		return
	}
}
