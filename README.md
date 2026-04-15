# To-Do
Менеджер для командной строки
👨
# Автор
Сергеев А.В.

## Описание
Консольное приложение для управления задачами.
Позволяет:
  - добавлять задачи
  - просматривать список задач
  - отмечать задачи выполненными
  - удалять задачи
  - импортировать и экспортировать задачи (JSON/CSV)

Основное хранилище данных — файл `tasks.json`.

## Требования
   - Go версии 1.16+
   - Терминал

## Запуск
Из корня проекта:
    ```bash
       go run ./cmd/todo <команда> [параметры]
    ```

# Команды:

## Добавить задачу
    go run ./cmd/todo add --desc="Купить продукты"
## Посмотреть задачи
    go run ./cmd/todo list
## Фильтрация:
    go run ./cmd/todo list --filter=all
    go run ./cmd/todo list --filter=done
    go run ./cmd/todo list --filter=pending
## Отметить задачу выполненной
    go run ./cmd/todo complete --id=1
## Удалить задачу
    go run ./cmd/todo delete --id=1

## Экспорт задач
### В JSON:
    go run ./cmd/todo export --format=json --out=tasks.json
### В CSV:
    go run ./cmd/todo export --format=csv --out=tasks.csv

## Импорт задач
### Из JSON:
    go run ./cmd/todo load --file=tasks.json
### Из CSV:
    go run ./cmd/todo load --file=tasks.csv

## Хранение данных
Все задачи хранятся в файле:
    tasks.json
Пример:
    [
        {
            "id": 1,
            "description": "Купить продукты",
            "done": false
        }
    ]

# Структура проекта
    To-Do/
        ├── cmd/
        │ └── todo/
        │ ├── main.go               // точка входа: парсинг аргументов, запуск команд
        ├── internal/
        │ ├── todo/
        │ │ ├── task.go             // модель Task и методы
        │ │ └── manager.go          // бизнес-логика (Add, List, Complete, Delete)
        │ └── storage/
        │ ├── json_storage.go       // функции LoadJSON, SaveJSON
        │ └── csv_storage.go        // функции LoadCSV, SaveCSV
        ├── go.mod                  // модули Go
        └── README.md               // документация и примеры использования

# Особенности
    Основной файл хранения — tasks.json
    CSV используется только для импорта/экспорта
    Все команды обрабатывают ошибки (ввод, файлы, id)

# Тестирование
    Запуск тестов:
    go test ./...

# Реализовано
    add
    list (с фильтрацией)
    complete
    delete
    export (JSON/CSV)
    load (JSON/CSV)