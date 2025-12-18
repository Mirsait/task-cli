package main

import (
	"fmt"
	"strings"

	"github.com/Mirsait/task-cli/models"
)

const (
	reset  = "\033[0m"
	bold   = "\033[1m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	gray   = "\033[90m"
)

func Clear() {
	fmt.Print("\033[H\033[2J")
}

func top() {
	dash5 := strings.Repeat("─", 5)
	dash14 := strings.Repeat("─", 14)
	dash18 := strings.Repeat("─", 18)
	dash42 := strings.Repeat("─", 42)
	fmt.Printf(bold+blue+"%s%s%s%s%s%s", "┌", dash5, "┬", dash14, "┬", dash18)
	fmt.Printf("%s%s%s%s%s\n", "┬", dash18, "┬", dash42, "┐"+reset)
}

func bottom() {
	dash5 := strings.Repeat("─", 5)
	dash14 := strings.Repeat("─", 14)
	dash18 := strings.Repeat("─", 18)
	dash42 := strings.Repeat("─", 42)
	fmt.Printf(bold+blue+"%s%s%s%s%s%s", "└", dash5, "┴", dash14, "┴", dash18)
	fmt.Printf("%s%s%s%s%s\n", "┴", dash18, "┴", dash42, "┘"+reset)
}

func PrintTasks(tasks []Task) {
	fmt.Printf(bold + yellow + "task-cli\n" + reset)
	top()
	for _, t := range tasks {
		prettyCreatedDate := t.CreatedAt.Format("02-01-2006 15:04")
		prettyUpdatedDate := t.UpdatedAt.Format("02-01-2006 15:04")
		statusColor := gray
		switch t.Status {
		case models.Todo:
			statusColor = green
		case models.Done:
			statusColor = red
		}
		fmt.Printf("│ %003d │ %s%12s%s │ %s │ %s │ %-40s │ \n",
			t.Id, statusColor, t.Status, reset, prettyCreatedDate, prettyUpdatedDate, t.Description)
	}
	bottom()
}
