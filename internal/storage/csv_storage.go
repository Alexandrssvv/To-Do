package storage

import (
	"encoding/csv"
	"os"
	"strconv"
	"todo-app/internal/todo"
)

func LoadCSV(path string) ([]todo.Task, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) <= 1 {
		return []todo.Task{}, nil
	}

	var tasks []todo.Task

	for _, record := range records[1:] {
		if len(record) < 3 {
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		done, err := strconv.ParseBool(record[2])
		if err != nil {
			return nil, err
		}

		task := todo.Task{
			ID:          id,
			Description: record[1],
			Done:        done,
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func SaveCSV(path string, tasks []todo.Task) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"ID", "Description", "Done"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, task := range tasks {
		record := []string{
			strconv.Itoa(task.ID),
			task.Description,
			strconv.FormatBool(task.Done),
		}

		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
