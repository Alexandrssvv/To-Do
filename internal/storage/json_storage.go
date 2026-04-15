package storage

import (
	"encoding/json"
	"os"
	"todo-app/internal/todo"
)

func LoadJSON(path string) ([]todo.Task, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		emptyTasks := []todo.Task{}
		if err := SaveJSON(path, emptyTasks); err != nil {
			return nil, err
		}
		return emptyTasks, nil
	}
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var tasks []todo.Task
	if len(data) == 0 {
		return []todo.Task{}, nil
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveJSON(path string, tasks []todo.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
