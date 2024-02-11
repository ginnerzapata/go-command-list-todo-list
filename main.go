package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

var tasks = []Task{}

func addTask(name string) {
	id, err := gonanoid.New(3)
	if err != nil {
		fmt.Println("Error generating ID:", err)
	}
	task:= Task{ID: string(id), Name: name, Done: false}
	tasks = append(tasks, task)
	fmt.Printf("Added task: %s\n", name)
}

func removeTask (id string) {
	for i, task := range tasks {
		if task.ID==id {
			tasks = append(tasks[:i], tasks[i+1:]... )
			fmt.Println("Removed tasks:",id)
			return
		}
	}
	fmt.Println("Task not found:", id)
}

func markTaskDone(id string) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			fmt.Println("Marked task as done:", id)
			return
		}
	}
	fmt.Println("Task not found:", id)
}

func listTasks() {
	for _, task := range tasks {
		status:= "Pending"
		if task.Done {
			status= "Done"
		}
		fmt.Printf("ID: %s, Task: %s, Status: %s\n", task.ID, task.Name, status)
	}
}

func main() {
	scanner:= bufio.NewScanner(os.Stdin)

	for{
		fmt.Println(("\nCommands: add <task>, done <taskID>, remove <taskID>, list, quit"))
		fmt.Print("> ")
		scanner.Scan()
		input:=scanner.Text()

		if input == "quit" {
			break
		}

		parts:= strings.SplitN(input, " ", 2)
		command:=parts[0]

		switch command {
		case "add":
			if len(parts) != 2 {
				fmt.Println("Invalid add command. You need to specify a task name")
				continue
			}
			addTask(parts[1])

		case "done":
			if len(parts) != 2 {
				fmt.Println("Invalid command. You need to specify a taskID")
				continue
			}
			markTaskDone(parts[1])
		case "remove":
            if len(parts) != 2 {
                fmt.Println("Invalid remove command. You need to specify a task ID.")
                continue
            }
            removeTask(parts[1])
	
		case "list":
			listTasks()
		default:
			fmt.Println("Unknown command:", command)
		}
	}
}