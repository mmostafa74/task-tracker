package main

import (
	"encoding/json"
	"os"
)

const tasksFile = "tasks.json"

func LoadTasks() ([]Task, error) {
	data, err := os.ReadFile(tasksFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(tasksFile, data, 0644)
}

func NextID(tasks []Task) int {
	maxID := 0
	if len(tasks) == 0 {
		return 1
	}

	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}
