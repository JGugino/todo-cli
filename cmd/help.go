package cmd

type HelpCmd struct{
	Commands map[string]*Command
}

func (help *HelpCmd) ExecuteAction(cmd *Command){
	helpDisplay := MakeHelpDisplay("Todo List CLI - Help Screen", help.Commands)
	ShowHelpDisplay(helpDisplay)
}