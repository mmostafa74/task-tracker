package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func printUsage(w io.Writer) {
	fmt.Fprint(w, `Usage: task <command> [<args>]

Flags:
  -h, --help         Show this help message
  -i, --interactive  Run in interactive mode

Commands:
  add <description>          Add a new task
  update <id> <description>  Update an existing task
  delete <id>                Delete a task
  list [status]              List all tasks (optional filter: todo, in-progress, done)
  mark-in-progress <id>      Mark a task as in-progress
  mark-done <id>             Mark a task as done
  get <id>                   Get a task by ID
  quit/exit/q                Exit the program in interactive mode
`)
}

func runCommand(command string, args []string) error {
	switch command {
	case "add":
		if len(args) < 1 {
			return fmt.Errorf("Usage: task add <description>")
		}
		description := strings.TrimSpace(args[0])
		id, err := CmdAdd(description)
		if err != nil {
			return err
		}
		fmt.Printf("Task added successfully (ID: %d).\n", id)
		return nil

	case "update":
		if len(args) < 2 {
			return fmt.Errorf("Usage: task update <id> <description>")
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("Invalid task ID.")
		}
		description := strings.TrimSpace(args[1])
		id, err = CmdUpdate(id, description)
		if err != nil {
			return err
		}
		fmt.Printf("Task updated successfully (ID: %d).\n", id)
		return nil

	case "delete":
		if len(args) < 1 {
			return fmt.Errorf("Usage: task delete <id>")
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("Invalid task ID.")
		}
		if err := CmdDelete(id); err != nil {
			return err
		}
		fmt.Println("Task deleted successfully.")
		return nil

	case "list":
		var status Status = ""
		if len(args) > 0 {
			if !validateStatus(Status(args[0])) {
				return fmt.Errorf("Invalid status. Valid statuses are: todo, in-progress, done")
			}
			status = Status(args[0])
		}
		tasks, err := CmdList(status)
		if err != nil {
			return err
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return nil
		}
		for _, task := range tasks {
			fmt.Println(task.String())
		}
		return nil

	case "mark-in-progress":
		if len(args) < 1 {
			return fmt.Errorf("Usage: task mark-in-progress <id>")
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("Invalid task ID.")
		}
		if err := CmdMark(id, StatusInProgress); err != nil {
			return err
		}
		fmt.Println("Task marked as in-progress successfully.")
		return nil

	case "mark-done":
		if len(args) < 1 {
			return fmt.Errorf("Usage: task mark-done <id>")
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("Invalid task ID.")
		}
		if err := CmdMark(id, StatusDone); err != nil {
			return err
		}
		fmt.Println("Task marked as done successfully.")
		return nil

	case "get":
		if len(args) < 1 {
			return fmt.Errorf("Usage: task get <id>")
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("Invalid task ID.")
		}
		task, err := CmdGetTask(id)
		if err != nil {
			return err
		}
		fmt.Println(task.String())
		return nil

	case "help", "-h", "--help":
		printUsage(os.Stdout)
		return nil

	case "quit", "exit", "q":
		return io.EOF

	default:
		return fmt.Errorf("Unknown command: %s. Try --help.", command)
	}
}

func parseLine(input string) ([]string, error) {
	var args []string
	var current strings.Builder
	inQuote := false
	for _, ch := range input {
		switch {
		case ch == '"':
			inQuote = !inQuote
		case ch == ' ' && !inQuote:
			if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}
		default:
			current.WriteRune(ch)
		}
	}
	if inQuote {
		return nil, fmt.Errorf("unclosed quote")
	}
	if current.Len() > 0 {
		args = append(args, current.String())
	}
	return args, nil
}

func runREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("task> ")
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			fmt.Print("task> ")
			continue
		}
		parts, err := parseLine(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			fmt.Print("task> ")
			continue
		}
		command := parts[0]
		args := parts[1:]

		err = runCommand(command, args)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Print("task> ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
