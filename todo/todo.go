package todo

import (
	"encoding/json"
	"errors"
	"os"
)

type Task struct {
	Name string
	Done bool
}

const filename = "todo/task.json"

func SaveTasks(tasks []Task) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)
}

func LoadTasks() ([]Task, error) {
	var tasks []Task

	file, err := os.Open(filename)
	if errors.Is(err, os.ErrNotExist) {
		// No tasks yet, return empty slice
		return tasks, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	return tasks, err
}

func AddTask(name string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}
	tasks = append(tasks, Task{Name: name})
	return SaveTasks(tasks)
}

func ListTasks() ([]Task, error) {
	return LoadTasks()
}

func CompleteTask(index int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}
	if index < 1 || index > len(tasks) {
		return errors.New("invalid task number")
	}
	tasks[index-1].Done = true
	return SaveTasks(tasks)
}