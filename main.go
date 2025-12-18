package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/Mirsait/task-cli/commands"
	"github.com/Mirsait/task-cli/models"
	"github.com/Mirsait/task-cli/storage"
)

type Task = models.Task
type Status = models.Status

var available_commands = []string{
	"add", "update", "delete", "list", "complete", "start", "help"}

func main() {
	Clear()
	args := os.Args[1:]
	if len(args) == 0 {
		log.Printf("Incorrect number of arguments")
		log.Println("Run `task-cli help` to see available commands.")
		return
	}
	cmd := args[0]

	if !slices.Contains(available_commands, cmd) {
		log.Printf("Undefinded command: %s\n", cmd)
		fmt.Println("Available commands: ", available_commands)
		return
	}

	if cmd == "help" {
		fmt.Println("Available commands:")
		fmt.Printf("%-32s - %s\n", "`add \"task description\"`", "add new task with description")
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

	loadData := func() ([]models.Task, error) {
		return storage.Load("tasks.json")
	}

	saveData := func(tasks []models.Task) error {
		return storage.Save("tasks.json", tasks)
	}

	switch cmd {
	case "add":
		if len(args) != 2 {
			log.Printf("Usage `task-cli add \"newDescription\"`\n")
			return
		}
		tasks, err := loadData()
		if err != nil {
			log.Printf("Cannot load tasks: %v\n", err)
			return
		}
		tasks, err = commands.Add(tasks, args[1])
		if err != nil {
			log.Printf("Cannot create task: %v\n", err)
			return
		}
		err = saveData(tasks)
		if err != nil {
			log.Printf("Cannot save tasks: %v\n", err)
		}
		log.Println("Task created.")
	case "delete":
		if len(args) != 2 {
			log.Println("Usage `task-cli delete taskId`")
			return
		}
		id, err := ParseToInt(args[1])
		if err != nil {
			log.Printf("Invalid task id: %v\n", err)
			return
		}
		tasks, err := loadData()
		if err != nil {
			log.Printf("Cannot load tasks: %v\n", err)
			return
		}
		tasks, err = commands.Delete(tasks, id)
		if err != nil {
			log.Printf("Cannot delete task: %v\n", err)
			return
		}
		err = saveData(tasks)
		if err != nil {
			log.Printf("Cannot save tasks: %v\n", err)
			return
		}
		log.Printf("Task %d deleted.\n", id)
	case "update":
		if len(args) != 3 {
			log.Println("Usage `task-cli update id \"new description\"`")
			return
		}
		id, err := ParseToInt(args[1])
		if err != nil {
			log.Printf("Incorrect task id: %v\n", err)
			return
		}
		tasks, err := loadData()
		if err != nil {
			log.Fatalf("Failed to load tasks: %v", err)
			return
		}
		newText := args[2]
		tasks, err = commands.Update(tasks, id, newText)
		if err != nil {
			log.Printf("Cannot update task: %v", err)
		}
		err = saveData(tasks)
		if err != nil {
			log.Printf("Cannot save tasks: %v", err)
		}
		log.Printf("Task %d updated.\n", id)
	case "start":
		if len(args) != 2 {
			log.Println("Usage `task-cli start id`")
			return
		}
		id, err := ParseToInt(args[1])
		if err != nil {
			log.Printf("Incorrect task id: %v\n", err)
			return
		}
		tasks, err := loadData()
		if err != nil {
			log.Printf("Cannot load taskss: %v", err)
			return
		}
		tasks, err = commands.Mark(tasks, id, models.Progress)
		if err != nil {
			log.Printf("Cannot mark task: %v", err)
			return
		}
		err = saveData(tasks)
		if err != nil {
			log.Printf("Cannot save tasks: %v", err)
		}
		log.Printf("Task %d started.\n", id)
	case "complete":
		if len(args) != 2 {
			log.Println("Usage `task-cli complete id`")
			return
		}
		id, err := ParseToInt(args[1])
		if err != nil {
			log.Printf("Incorrect task id: %v\n", err)
			return
		}
		tasks, err := loadData()
		if err != nil {
			log.Printf("Cannot load taskss: %v", err)
			return
		}
		tasks, err = commands.Mark(tasks, id, models.Done)
		if err != nil {
			log.Printf("Cannot mark task: %v", err)
			return
		}
		err = saveData(tasks)
		if err != nil {
			log.Printf("Cannot save tasks: %v", err)
		}
		log.Printf("Task %d completed.\n", id)
	case "list":
		tasks, err := loadData()
		if err != nil {
			log.Printf("Cannot load tasks: %v", err)
			return
		}
		filterFn := func(t Task) bool { return true }
		if len(args) == 2 {
			status := args[1]
			switch status {
			case "todo":
				filterFn = func(t Task) bool { return t.Status != models.Done }
			case "done":
				filterFn = func(t Task) bool { return t.Status == models.Done }
			case "in-progress":
				filterFn = func(t Task) bool { return t.Status == models.Progress }
			default:
				log.Println("Undefinded status: ", status)
			}
		}
		tasks = commands.Filter(tasks, filterFn)
		PrintTasks(tasks)
	default:
		log.Println("Undefinded command")
	}
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}

func PrintTasks(tasks []Task) {
	fmt.Printf("task list: %d\n", len(tasks))
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
