package cmd

import (
	"errors"
	"fmt"
)

type CommandAction interface{
	ExecuteAction(cmd *Command)
}

type CommandHandler struct{
	commands map[string]*Command
	args []string
	command string

}

type Command struct{
	CommandName string
	CommandDesc string
	CommandAction CommandAction
}

type HelpDisplay struct {
	displayTitle string
	displayCmds map[string]string
}

func NewCommandHandler(args []string, command string) *CommandHandler{
	return &CommandHandler{
		commands: make(map[string]*Command, 6),
		args: args,
		command: command,
	}
}

func MakeHelpDisplay(displayTitle string, commands map[string]*Command) HelpDisplay{

	convertedCmds := make(map[string]string, 6);

	for _, c := range commands {
		convertedCmds[c.CommandName] = c.CommandDesc
	}

	return HelpDisplay{
		displayTitle: displayTitle,
		displayCmds: convertedCmds,
	}
}

func ShowHelpDisplay(helpDisplay HelpDisplay){
	fmt.Println(helpDisplay.displayTitle)
	fmt.Println("----------------------------")
	for n, v := range helpDisplay.displayCmds {
		fmt.Printf("Command: %s - %s \n", n, v)
	}
}

func (handler *CommandHandler) GetCommands() map[string]*Command{
	return handler.commands
}

func (handler *CommandHandler) AddCommand(command *Command, ){
	handler.commands[command.CommandName] = command
}

func (handler *CommandHandler) ValidateAndReturnCommand() (*Command, error){
	found := handler.commands[handler.command]

	if found == nil{
		return nil, errors.New("no command found");
	}

	return found, nil
}

func (handler *CommandHandler) HandleCommand(command *Command){
	command.CommandAction.ExecuteAction(command)
}