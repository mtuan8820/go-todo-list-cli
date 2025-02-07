package main

import (
	"go-todo-list-cli/pkg/commands"
	"go-todo-list-cli/pkg/models"
	"log"
)

func main() {
	todoList := models.TodoList{}

	err := todoList.Load("todo.json")
	if err != nil {
		log.Fatal(err)
	}
	cmdFlags := commands.NewCmdFlag()
	cmdFlags.Execute(&todoList)
	todoList.Save("todo.json")
}
