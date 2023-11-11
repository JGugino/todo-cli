package cmd

import (
	"github.com/JGugino/todo-cli/todo"
)

type CompleteCmd struct{
	TodoHandler *todo.TodoHandler
}

func (comp *CompleteCmd) ExecuteAction(handler *CommandHandler, cmd *Command){
	// todoName, err := handler.ValidateAndReturnFlag("-n")

	// if err != nil{
	// 	panic("Invalid arguments, type 'todo help' for more information.")
	// }

	// add.TodoHandler.AddTodo("Default List", todoName, todoValue, false)
	// handler.MustSaveTodoListAfterAction(*add.TodoHandler.TodoLists["Default List"])

	// fmt.Println("Todo item added successfully")
}