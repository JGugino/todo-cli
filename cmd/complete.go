package cmd

import (
	"fmt"

	"github.com/JGugino/todo-cli/todo"
)

type CompleteCmd struct{
	TodoHandler *todo.TodoHandler
}

func (comp *CompleteCmd) ExecuteAction(handler *CommandHandler, cmd *Command){
	todoName, err := handler.ValidateAndReturnFlag("-n")

	if err != nil{
		panic("Invalid arguments, type 'todo help' for more information.")
	}

	compErr := comp.TodoHandler.CompleteTodo("Default List", todoName)

	if compErr != nil {
		panic(compErr)
	}

	handler.MustSaveTodoListAfterAction(*comp.TodoHandler.TodoLists["Default List"])

	fmt.Printf("Todo (%s) has been completed. \n", todoName)
}