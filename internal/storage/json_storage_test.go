package storage

import (
	"path/filepath"
	"testing"
	"todo-app/internal/todo"
)

func TestSaveAndLoadJSON(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "tasks.json")

	expectedTasks := []todo.Task{
		{ID: 1, Description: "Тест 1", Done: false},
		{ID: 2, Description: "Тест 2", Done: true},
	}

	err := SaveJSON(filePath, expectedTasks)
	if err != nil {
		t.Fatalf("ошибка при сохранении JSON: %v", err)
	}

	loadedTasks, err := LoadJSON(filePath)
	if err != nil {
		t.Fatalf("ошибка при загрузке JSON: %v", err)
	}

	if len(loadedTasks) != len(expectedTasks) {
		t.Fatalf("ожидалось %d задач, получено %d", len(expectedTasks), len(loadedTasks))
	}

	for i := range expectedTasks {
		if loadedTasks[i] != expectedTasks[i] {
			t.Errorf("ожидалось %+v, получено %+v", expectedTasks[i], loadedTasks[i])
		}
	}
}
