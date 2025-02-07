package commands

import (
	"flag"
	"fmt"
	"go-todo-list-cli/pkg/models"
)

type CmdFlags struct {
	Add    string
	Delete int
	Toggle int
	Print  bool
}

func NewCmdFlag() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "Add", "", "Add a new todo task's title")
	flag.IntVar(&cf.Delete, "Delete", -1, "Delete a todo task at index position")
	flag.IntVar(&cf.Toggle, "Toggle", -1, "Toggle state (completed) of a task at index position")
	flag.BoolVar(&cf.Print, "Print", false, "Print a table show all tasks")

	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute(todoList *models.TodoList) {
	switch {
	case cf.Print:
		todoList.Print()
	case cf.Add != "":
		todoList.Add(cf.Add)
	case cf.Delete != -1:
		todoList.Delete(cf.Delete)
	case cf.Toggle != -1:
		todoList.Toggle(cf.Toggle)
	default:
		fmt.Printf("Invalid command")
	}
}
