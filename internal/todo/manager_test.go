package todo

import "testing"

func TestAdd(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Первая задача", Done: false},
	}

	updatedTasks := Add(tasks, "Новая задача")

	if len(updatedTasks) != 2 {
		t.Fatalf("ожидалось 2 задачи, получено %d", len(updatedTasks))
	}

	lastTask := updatedTasks[1]

	if lastTask.ID != 2 {
		t.Errorf("ожидался ID = 2, получено %d", lastTask.ID)
	}

	if lastTask.Description != "Новая задача" {
		t.Errorf("ожидалось описание %q, получено %q", "Новая задача", lastTask.Description)
	}

	if lastTask.Done != false {
		t.Errorf("ожидалось Done = false, получено %v", lastTask.Done)
	}
}

func TestListAll(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Задача 1", Done: false},
		{ID: 2, Description: "Задача 2", Done: true},
	}

	result := List(tasks, "all")

	if len(result) != 2 {
		t.Errorf("ожидалось 2 задачи, получено %d", len(result))
	}
}

func TestListDone(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Задача 1", Done: false},
		{ID: 2, Description: "Задача 2", Done: true},
		{ID: 3, Description: "Задача 3", Done: true},
	}

	result := List(tasks, "done")

	if len(result) != 2 {
		t.Fatalf("ожидалось 2 выполненные задачи, получено %d", len(result))
	}

	for _, task := range result {
		if !task.Done {
			t.Errorf("ожидалась выполненная задача, получено Done = false")
		}
	}
}

func TestListPending(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Задача 1", Done: false},
		{ID: 2, Description: "Задача 2", Done: true},
		{ID: 3, Description: "Задача 3", Done: false},
	}

	result := List(tasks, "pending")

	if len(result) != 2 {
		t.Fatalf("ожидалось 2 невыполненные задачи, получено %d", len(result))
	}

	for _, task := range result {
		if task.Done {
			t.Errorf("ожидалась невыполненная задача, получено Done = true")
		}
	}
}

func TestCompleteSuccess(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Задача 1", Done: false},
		{ID: 2, Description: "Задача 2", Done: false},
	}

	updatedTasks, err := Complete(tasks, 2)
	if err != nil {
		t.Fatalf("не ожидалась ошибка, получено: %v", err)
	}

	if !updatedTasks[1].Done {
		t.Errorf("ожидалось, что задача будет выполненной")
	}
}

func TestCompleteNotFound(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Задача 1", Done: false},
	}

	_, err := Complete(tasks, 99)
	if err == nil {
		t.Fatal("ожидалась ошибка, но получили nil")
	}
}

func TestDeleteSuccess(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Задача 1", Done: false},
		{ID: 2, Description: "Задача 2", Done: true},
	}

	updatedTasks, err := Delete(tasks, 1)
	if err != nil {
		t.Fatalf("не ожидалась ошибка, получено: %v", err)
	}

	if len(updatedTasks) != 1 {
		t.Fatalf("ожидалась 1 задача, получено %d", len(updatedTasks))
	}

	if updatedTasks[0].ID != 2 {
		t.Errorf("ожидалось, что останется задача с ID = 2, получено %d", updatedTasks[0].ID)
	}
}

func TestDeleteNotFound(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Задача 1", Done: false},
	}

	_, err := Delete(tasks, 99)
	if err == nil {
		t.Fatal("ожидалась ошибка, но получили nil")
	}
}
