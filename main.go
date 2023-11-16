package main

import (
	"fmt"
	"os"

	"github.com/JGugino/todo-cli/cmd"
	"github.com/JGugino/todo-cli/todo"
)

func main(){
	if len(os.Args) <= 1{
		fmt.Println("Invalid command! type 'todo help' for more information.")
		return
	}

	commandHandler := cmd.NewCommandHandler(os.Args[2:], os.Args[1:2][0])
	todoHandler := todo.NewTodoHandler()

	list, err := todo.MustLoadTodoList()

	if err != nil {
		panic("Unable to load todo lists, please try again.")
	}

	todoHandler.AddList(list.ListName, make(map[string]todo.TodoItem))

	for n, v := range list.TodoItems {
		todoHandler.AddTodo(list.ListName, n, v.TodoValue, v.Completed)	
	}

	commandHandler.AddCommand(&cmd.Command{CommandName: "add", CommandDesc: "Adds a new todo item to the list. (-n='NAME OF TODO' -v='VALUE OF TODO')", CommandAction: &cmd.AddCmd{TodoHandler: todoHandler}})
	commandHandler.AddCommand(&cmd.Command{CommandName: "remove", CommandDesc: "Removes a command from the todo list (-n='NAME OF TODO')", CommandAction: &cmd.RemoveCmd{TodoHandler: todoHandler}})
	commandHandler.AddCommand(&cmd.Command{CommandName: "complete", CommandDesc: "Marks the specified todo as completed (-n='NAME OF TODO')", CommandAction: &cmd.CompleteCmd{TodoHandler: todoHandler}})
	commandHandler.AddCommand(&cmd.Command{CommandName: "list", CommandDesc: "Lists out all items on the todo list", CommandAction: &cmd.ListCmd{TodoHandler: todoHandler}})
	commandHandler.AddCommand(&cmd.Command{CommandName: "help", CommandDesc: "Prints out a list of all available commands and their descriptions", CommandAction: &cmd.HelpCmd{Commands: commandHandler.GetCommands()}})

	currentCmd, err := commandHandler.ValidateAndReturnCommand();

	if err != nil{
		panic(err)
	}

	commandHandler.HandleCommand(commandHandler, currentCmd)
}