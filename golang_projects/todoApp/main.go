package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID           uuid.UUID `json:"id"`
	CreationDate time.Time `json:"creation_date"`
	Task         string    `json:"task"`
	State        bool      `json:"state"`
}

const JSON_PATH = "todo.json"

func main() {
	tasks := []Task{}

	t := Task{
		ID:           uuid.New(), // генерируем новый уникальный ID
		CreationDate: time.Now(),
		Task:         "Сделать проект",
		State:        false,
	}

	tasks = append(tasks, t)
	fmt.Print(tasks)

}
