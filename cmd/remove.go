package cmd

import (
	"fmt"

	"github.com/JGugino/todo-cli/todo"
)

type RemoveCmd struct{
	TodoHandler *todo.TodoHandler
}

func (remove *RemoveCmd) ExecuteAction(handler *CommandHandler, cmd *Command){
	todoName, err := handler.ValidateAndReturnFlag("-n")

	if err != nil{
		panic("Invalid arguments, type 'todo help' for more information.")
	}

	remErr := remove.TodoHandler.RemoveTodo("Default List", todoName)
	
	if remErr != nil {
		panic(remErr)
	}

	handler.MustSaveTodoListAfterAction(*remove.TodoHandler.TodoLists["Default List"])

	fmt.Printf("Todo (%s) has been removed. \n", todoName)
}