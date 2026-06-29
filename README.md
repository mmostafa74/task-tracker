# Task Tracker CLI

A command-line task tracker written in Go. Tasks are stored as JSON in `tasks.json` — zero external dependencies.

## Commands

| Command | Usage | Description |
|---|---|---|
| `add` | `task add "<description>"` | Create a new task (status: todo) |
| `update` | `task update <id> "<description>"` | Update a task's description |
| `delete` | `task delete <id>` | Remove a task |
| `mark-in-progress` | `task mark-in-progress <id>` | Mark a task as in-progress |
| `mark-done` | `task mark-done <id>` | Mark a task as done |
| `list` | `task list [status]` | List all tasks. Optional filter: `todo`, `in-progress`, or `done` |
| `--help` | `task --help` | Show usage information |

## Quick Start

```bash
# Add a task
go run . add "Buy groceries"

# List all tasks
go run . list

# Mark a task as done
go run . mark-done 1

# List only done tasks
go run . list done

# Update a task
go run . update 1 "Buy groceries and cook dinner"

# Delete a task
go run . delete 1
```

## Build

```bash
go build -o task-cli
./task-cli add "Build the binary"
```

## Project Structure

``` bash
task-tracker/
├── main.go       # CLI entry point — arg parsing and routing
├── task.go       # Task struct, status constants, String() with colors
├── storage.go    # Load/Save tasks.json, ID generation
├── commands.go   # Business logic for each command
├── tasks.json    # Data file (created at runtime)
└── go.mod        # Module definition
```
