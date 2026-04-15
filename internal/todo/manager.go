package todo

import "fmt"

func Add(tasks []Task, desc string) []Task {
	newTask := Task{
		ID:          nextID(tasks),
		Description: desc,
		Done:        false,
	}

	return append(tasks, newTask)
}

func List(tasks []Task, filter string) []Task {
	var result []Task

	for _, task := range tasks {
		switch filter {
		case "all":
			result = append(result, task)
		case "done":
			if task.Done {
				result = append(result, task)
			}
		case "pending":
			if !task.Done {
				result = append(result, task)
			}
		}
	}

	return result
}

func Complete(tasks []Task, id int) ([]Task, error) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			return tasks, nil
		}
	}

	return tasks, fmt.Errorf("Задача с id %d не найдена", id)
}

func Delete(tasks []Task, id int) ([]Task, error) {
	for i := range tasks {
		if tasks[i].ID == id {
			return append(tasks[:i], tasks[i+1:]...), nil
		}
	}

	return tasks, fmt.Errorf("Задача с id %d не найдена", id)
}

func nextID(tasks []Task) int {
	maxID := 0

	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	return maxID + 1
}
