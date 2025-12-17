package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var commands = []string{"add", "update", "delete", "list", "complete", "start", "help"}

func main() {
	Clear()
	args := os.Args[1:]
	if len(args) == 0 {
		log.Printf("Incorrect number of arguments")
		log.Println("Run `task-cli help` to see available commands.")
		return
	}
	command := args[0]

	if !slices.Contains(commands, command) {
		log.Printf("Undefinded command: %s\n", command)
		fmt.Println("Available commands: ", commands)
		return
	}

	if command == "help" {
		fmt.Println("Available commands:")
		fmt.Printf("%-32s - %s\n", "`add \"description\"`", "add new task with description")
		fmt.Printf("%-32s - %s\n", "`delete id`", "delete task by id")
		fmt.Printf("%-32s - %s\n", "`update id new-description`", "update task's description")
		fmt.Printf("%-32s - %s\n", "`complete id`", "mark task as completed")
		fmt.Printf("%-32s - %s\n", "`start id`", "mark task as started")
		fmt.Printf("%-32s - %s\n", "`list`", "list of all tasks")
		fmt.Printf("%-32s - %s\n", "`list done`", "list of all completed tasks")
		fmt.Printf("%-32s - %s\n", "`list todo`", "list of all outstanding tasks")
		fmt.Printf("%-32s - %s\n", "`list in-progress`", "list of all tasks that are in progress")
		return
	}

	repo := Repository{filename: "tasks.json"}
	repo.load()

	switch command {
	case "add":
		if len(args) != 2 {
			log.Printf("Incorrect number of arguments; expected `add \"newDescription\"`\n")
			return
		}
		repo.add(args[1])
	case "delete":
		if len(args) != 2 {
			log.Printf("Incorrect number of arguments; expected `delete taskId` \n")
			return
		}
		id, err := ParseToInt(args[1])
		if err != nil {
			log.Printf("Error parsing the task's id: %s\n", args[1])
			return
		}
		repo.delete(id)
	case "update":
		id, err := ParseToInt(args[1])
		if err != nil {
			log.Printf("Error parsing the task's id: %s\n", args[1])
			return
		}
		newText := args[2]
		if newText == "" {
			log.Println("New task's description is empty.")
			return
		}
		repo.update(id, newText)
	case "start":
		id, err := ParseToInt(args[1])
		if err != nil {
			log.Printf("Error parsing the task's id: %s\n", args[1])
			return
		}
		repo.mark(id, Progress)
	case "complete":
		id, err := ParseToInt(args[1])
		if err != nil {
			log.Printf("Error parsing the task's id: %s\n", args[1])
			return
		}
		repo.mark(id, Done)
	case "list":
		if len(args) == 2 {
			filter := args[1]
			switch filter {
			case "todo":
				tasks := repo.filter(func(t Task) bool { return t.Status != Done })
				PrintTasks(tasks)
			case "done":
				tasks := repo.filter(func(t Task) bool { return t.Status == Done })
				PrintTasks(tasks)
			case "in-progress":
				tasks := repo.filter(func(t Task) bool { return t.Status == Progress })
				PrintTasks(tasks)
			default:
				log.Println("Undefinded status: ", args[1])
			}
		} else {
			tasks := repo.getAll()
			PrintTasks(tasks)
		}
	default:
		log.Println("Undefinded command")
	}
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}

func PrintTasks(tasks []Task) {
	for _, t := range tasks {
		prettyCreatedDate := t.CreatedAt.Format("02-01-2006 15:04")
		prettyUpdatedDate := t.UpdatedAt.Format("02-01-2006 15:04")
		fmt.Printf("%03d | %12s | %s | %s | %s\n",
			t.Id, t.Status, prettyCreatedDate, prettyUpdatedDate, t.Description)
	}
}

func ParseToInt(value string) (int, error) {
	trim := strings.TrimSpace(value)
	num, err := strconv.Atoi(trim)
	if err != nil {
		return 0, err
	}
	return num, nil
}
