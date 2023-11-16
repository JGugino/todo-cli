package cmd

import (
	"fmt"

	"github.com/JGugino/todo-cli/todo"
)

type AddCmd struct{
	TodoHandler *todo.TodoHandler
}

func (add *AddCmd) ExecuteAction(handler *CommandHandler, cmd *Command){
	todoName, err1 := handler.ValidateAndReturnFlag("-n")
	todoValue, err2 := handler.ValidateAndReturnFlag("-v")

	if err1 != nil || err2 != nil{
		panic("Invalid arguments, type 'todo help' for more information.")
	}

	addErr := add.TodoHandler.AddTodo("Default List", todoName, todoValue, false)

	if addErr != nil {
		panic(addErr)
	}

	handler.MustSaveTodoListAfterAction(*add.TodoHandler.TodoLists["Default List"])

	fmt.Printf("Todo (%s) has been added. \n", todoName)
}