package main

import (
	"errors"
	"time"
)

func validateStatus(status Status) bool {
	return status == StatusTodo || status == StatusInProgress || status == StatusDone
}

func CmdAdd(description string) (int, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return 0, err
	}

	newTask := Task{
		ID:          NextID(tasks),
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	err = SaveTasks(tasks)
	if err != nil {
		return 0, err
	}

	return newTask.ID, nil
}

func CmdUpdate(id int, description string) (int, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return 0, err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()

			err = SaveTasks(tasks)
			if err != nil {
				return 0, err
			}
			return tasks[i].ID, nil
		}
	}
	return 0, errors.New("task not found")
}

func CmdDelete(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)

			err = SaveTasks(tasks)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("task not found")
}

func CmdMark(id int, status Status) error {
	if !validateStatus(status) {
		return errors.New("invalid status")
	}
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()

			err = SaveTasks(tasks)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("task not found")
}

func CmdList(status Status) ([]Task, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return nil, err
	}

	if status == "" {
		return tasks, nil
	}
	if !validateStatus(status) {
		return nil, errors.New("invalid status")
	}
	if len(tasks) == 0 {
		return []Task{}, nil
	}
	var filteredTasks []Task
	for _, task := range tasks {
		if task.Status == status {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks, nil
}

func CmdGetTask(id int) (*Task, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return nil, err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}
