package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"todo-app/internal/storage"
	"todo-app/internal/todo"
)

const dataFile = "tasks.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("команда не указана")
		fmt.Println("пример использования:")
		fmt.Println("todo add --desc=\"Купить продукты\"")
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	tasks, err := storage.LoadJSON(dataFile)
	if err != nil {
		fmt.Println("Ошибка загрузки задач:", err)
		os.Exit(1)
	}

	switch cmd {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		desc := addCmd.String("desc", "", "Описание задачи")

		err = addCmd.Parse(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if *desc == "" {
			fmt.Println("Описание задачи не может быть пустым")
			os.Exit(1)
		}

		tasks = todo.Add(tasks, *desc)

		err = storage.SaveJSON(dataFile, tasks)
		if err != nil {
			fmt.Println("Ошибка сохранения задач:", err)
			os.Exit(1)
		}

		fmt.Println("Задача добавлена")

	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		filter := listCmd.String("filter", "all", "all, done, pending")

		err = listCmd.Parse(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if *filter != "all" && *filter != "done" && *filter != "pending" {
			fmt.Println("Неверный фильтр, используйте: all, done, pending")
			os.Exit(1)
		}

		filteredTasks := todo.List(tasks, *filter)

		if len(filteredTasks) == 0 {
			fmt.Println("Задачи не найдены")
			return
		}

		for _, task := range filteredTasks {
			status := " "
			if task.Done {
				status = "x"
			}

			fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Description)
		}

	case "complete":
		completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
		id := completeCmd.Int("id", 0, "id задачи")

		err = completeCmd.Parse(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if *id <= 0 {
			fmt.Println("id должен быть больше 0")
			os.Exit(1)
		}

		tasks, err = todo.Complete(tasks, *id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = storage.SaveJSON(dataFile, tasks)
		if err != nil {
			fmt.Println("Ошибка сохранения задач:", err)
			os.Exit(1)
		}

		fmt.Println("Задача отмечена как выполненная")

	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := deleteCmd.Int("id", 0, "id задачи")

		err = deleteCmd.Parse(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if *id <= 0 {
			fmt.Println("id должен быть больше 0")
			os.Exit(1)
		}

		tasks, err = todo.Delete(tasks, *id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = storage.SaveJSON(dataFile, tasks)
		if err != nil {
			fmt.Println("Ошибка сохранения задач:", err)
			os.Exit(1)
		}

		fmt.Println("Задача удалена")

	case "load":
		loadCmd := flag.NewFlagSet("load", flag.ExitOnError)
		file := loadCmd.String("file", "", "Путь к файлу")

		err = loadCmd.Parse(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if *file == "" {
			fmt.Println("Путь к файлу не может быть пустым")
			os.Exit(1)
		}

		ext := filepath.Ext(*file)

		switch ext {
		case ".json":
			tasks, err = storage.LoadJSON(*file)
		case ".csv":
			tasks, err = storage.LoadCSV(*file)
		default:
			fmt.Println("Неподдерживаемый формат файла, используйте .json или .csv")
			os.Exit(1)
		}

		if err != nil {
			fmt.Println("Ошибка загрузки файла:", err)
			os.Exit(1)
		}

		err = storage.SaveJSON(dataFile, tasks)
		if err != nil {
			fmt.Println("Ошибка сохранения задач:", err)
			os.Exit(1)
		}

		fmt.Println("tasks loaded")

	case "export":
		exportCmd := flag.NewFlagSet("export", flag.ExitOnError)
		format := exportCmd.String("format", "", "json или csv")
		out := exportCmd.String("out", "", "Путь к выходному файлу")

		err = exportCmd.Parse(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if *format == "" || *out == "" {
			fmt.Println("Необходимо указать формат и путь к файлу")
			os.Exit(1)
		}

		switch *format {
		case "json":
			err = storage.SaveJSON(*out, tasks)
		case "csv":
			err = storage.SaveCSV(*out, tasks)
		default:
			fmt.Println("Неподдерживаемый формат экспорта, используйте json или csv")
			os.Exit(1)
		}

		if err != nil {
			fmt.Println("Ошибка экспорта задач:", err)
			os.Exit(1)
		}

		fmt.Println("Задачи экспортированы")

	default:
		fmt.Println("Неизвестная команда", cmd)
		fmt.Println("Доступные команды: add, list, complete, delete, load, export")
		os.Exit(1)
	}
}
