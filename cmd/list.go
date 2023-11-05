package cmd

import (
	"fmt"

	"github.com/JGugino/todo-cli/todo"
)

type ListCmd struct{
	TodoHandler *todo.TodoHandler
	Commands map[string]*Command
}

func (list *ListCmd) ExecuteAction(cmd *Command){
	fmt.Printf("Active Todo Items (%s): \n", list.TodoHandler.TodoLists["Default List"].ListName)
	fmt.Println("------------------------------")
	for n, v := range list.TodoHandler.ListTodoItemsFromSpecifiedList("Default List") {
		if v.Completed {
			fmt.Printf("%s: [X] \n", n)
		}else{
			fmt.Printf("%s: [] \n", n)
		}
		fmt.Println(v.TodoValue)
	}
}