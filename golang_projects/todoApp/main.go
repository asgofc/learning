package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Тип для списка задач
type Todos []string

// Универсальное хранилище (generic)
type Storage[T any] struct {
	filename string
}

func NewStorage[T any](filename string) *Storage[T] {
	return &Storage[T]{filename: filename}
}

// Загрузка из файла
func (s *Storage[T]) Load(data *T) error {
	file, err := os.Open(s.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(data)
}

// Сохранение в файл
func (s *Storage[T]) Save(data *T) error {
	file, err := os.Create(s.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(data)
}

func main() {
	// Инициализируем пустой список задач
	todos := Todos{}

	// Создаем сторедж, привязанный к файлу todos.json
	storage := NewStorage[Todos]("todos.json")

	// Пробуем загрузить задачи
	err := storage.Load(&todos)
	if err != nil {
		fmt.Println("Warning: Could not load todos from storage. Starting fresh todos.")
	}

	// Добавим тестовую задачу и сохраним
	todos = append(todos, "Выучить Go")
	_ = storage.Save(&todos)

	fmt.Println("Todos:", todos)
}
