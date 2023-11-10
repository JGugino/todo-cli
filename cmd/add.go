package cmd

import (
	"fmt"

	"github.com/JGugino/todo-cli/todo"
)

type AddCmd struct{
	TodoHandler *todo.TodoHandler
}

func (add *AddCmd) ExecuteAction(handler *CommandHandler, cmd *Command){
	todoName, err := handler.ValidateAndReturnFlag("-n")
	todoValue, err := handler.ValidateAndReturnFlag("-v")

	if err != nil{
		panic("Invalid arguments, type 'todo help' for more information.")
	}

	add.TodoHandler.AddTodo("Default List", todoName, todoValue, false)
	handler.MustSaveTodoListAfterAction(*add.TodoHandler.TodoLists["Default List"])

	fmt.Println("Todo item added successfully")
}