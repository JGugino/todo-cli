package main

import (
	"os"

	"github.com/JGugino/todo-cli/cmd"
	"github.com/JGugino/todo-cli/todo"
)

func main(){
	if len(os.Args) <= 1{
		panic("Not enough arguments")
	}

	commandHandler := cmd.NewCommandHandler(os.Args[2:], os.Args[1:2][0])
	todoHandler := todo.NewTodoHandler()

	todoHandler.AddList("Default List", make(map[string]todo.TodoItem))
	todoHandler.AddTodo("Default List", "Exmaple Todo", "This is an example todo item", false)
	todoHandler.AddTodo("Default List", "Exmaple Todo 2", "This is an example todo item", true)
	todoHandler.AddTodo("Default List", "Exmaple Todo 3", "This is an example todo item", false)

	commandHandler.AddCommand(&cmd.Command{CommandName: "add", CommandDesc: "Adds a new todo item to the list.", CommandAction: &cmd.AddCmd{TodoHandler: todoHandler}})
	commandHandler.AddCommand(&cmd.Command{CommandName: "remove", CommandDesc: "Removes a command from the todo list", CommandAction: &cmd.RemoveCmd{TodoHandler: todoHandler}})
	commandHandler.AddCommand(&cmd.Command{CommandName: "list", CommandDesc: "Lists out all items on the todo list", CommandAction: &cmd.ListCmd{TodoHandler: todoHandler}})
	commandHandler.AddCommand(&cmd.Command{CommandName: "help", CommandDesc: "Prints out a list of all available commands and their descriptions", CommandAction: &cmd.HelpCmd{Commands: commandHandler.GetCommands()}})

	currentCmd, err := commandHandler.ValidateAndReturnCommand();

	if err != nil{
		panic(err)
	}

	commandHandler.HandleCommand(currentCmd)
}