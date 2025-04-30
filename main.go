package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/amiothenes/go-todo/todo"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Expected 'add', 'list' or 'done' subcommands")
		return
	}

	switch args[0] {
	case "add":
		if len(args) < 2 {
			fmt.Println("Usage: go-todo add <task>")
			return
		}
		task := args[1]
		if err := todo.AddTask(task); err != nil {
			fmt.Println("Error adding task:", err)
		} else {
			fmt.Println("Task added:", task)
		}

	case "list":
		tasks, err := todo.ListTasks()
		if err != nil {
			fmt.Println("Error listing tasks:", err)
			return
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks found!")
			return
		}
		for i, t := range tasks {
			status := " "
			if t.Done {
				status = "✔"
			}
			fmt.Printf("%d. [%s] %s\n", i+1, status, t.Name)
		}

	case "done":
		if len(args) < 2 {
			fmt.Println("Usage: go-todo done <task number>")
			return
		}
		i, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid number:", args[1])
			return
		}
		if err := todo.CompleteTask(i); err != nil {
			fmt.Println("Error completing task:", err)
		} else {
			fmt.Printf("Marked task %d as done.\n", i)
		}

	default:
		fmt.Println("Unknown command:", args[0])
	}
}