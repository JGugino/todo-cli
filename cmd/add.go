package cmd

import (
	"github.com/JGugino/todo-cli/todo"
)

type AddCmd struct{
	TodoHandler *todo.TodoHandler
}

func (add *AddCmd) ExecuteAction(cmd *Command){
	
}