package cmd

import (
	"fmt"

	"github.com/JGugino/todo-cli/todo"
)

type RemoveCmd struct{
	TodoHandler *todo.TodoHandler
}

func (remove *RemoveCmd) ExecuteAction(handler *CommandHandler, cmd *Command){
	fmt.Println("Removing todo")
}