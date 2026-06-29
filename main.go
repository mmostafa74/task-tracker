package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func printUsage(w io.Writer) {
	fmt.Fprintln(w, "Usage: task <command> [<args>]")
	fmt.Fprintln(w, "Commands:")
	fmt.Fprintln(w, "  add <description>       Add a new task")
	fmt.Fprintln(w, "  update <id>(int) <description> Update an existing task")
	fmt.Fprintln(w, "  delete <id>(int)             Delete a task")
	fmt.Fprintln(w, "  list  <status>(optional)          List all tasks")
	fmt.Fprintln(w, "  mark-in-progress <id>(int) Mark a task as in-progress")
	fmt.Fprintln(w, "  mark-done <id>(int)       Mark a task as done")
}

func main() {
	if len(os.Args) < 2 {
		printUsage(os.Stderr)
		os.Exit(1)
	}

	command := os.Args[1]
	if command == "--help" {
		printUsage(os.Stdout)
		os.Exit(0)
	}
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Usage: task add <description>")
			os.Exit(1)
		}
		description := strings.TrimSpace(os.Args[2])
		err := CmdAdd(description)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Task added successfully.")
	case "update":
		if len(os.Args) < 4 {
			fmt.Fprintln(os.Stderr, "Usage: task update <id>(int) <description>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid task ID.")
			os.Exit(1)
		}
		description := strings.TrimSpace(os.Args[3])
		err = CmdUpdate(id, description)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Task updated successfully.")
	case "delete":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Usage: task delete <id>(int)")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid task ID.")
			os.Exit(1)
		}
		err = CmdDelete(id)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Task deleted successfully.")
	case "list":
		var status Status = ""
		if len(os.Args) > 2 {
			if !validateStatus(Status(os.Args[2])) {
				fmt.Fprintln(os.Stderr, "Invalid status. Valid statuses are: todo, in-progress, done")
				os.Exit(1)
			}
			status = Status(os.Args[2])
		}
		tasks, err := CmdList(status)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		for _, task := range tasks {
			fmt.Println(task.String())
		}
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Usage: task mark-in-progress <id>(int)")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid task ID.")
			os.Exit(1)
		}
		err = CmdMark(id, StatusInProgress)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Task marked as in-progress successfully.")
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Usage: task mark-done <id>(int)")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid task ID.")
			os.Exit(1)
		}
		err = CmdMark(id, StatusDone)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Task marked as done successfully.")
	default:
		fmt.Fprintln(os.Stderr, "Unknown command:", command)
		os.Exit(1)
	}
}
