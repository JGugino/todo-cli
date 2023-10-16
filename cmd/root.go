/*
Copyright © 2023 Joshua Gugino <gugino.inquires@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/JGugino/todo-cli/cmd/helpers"
	"github.com/spf13/cobra"
)

var PrimaryListHandler *helpers.ListHandler

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "A simple CLI Todo app",
	Long: `A simple and easy to use CLI Todo app to keep your schedule organized`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	PrimaryListHandler = &helpers.ListHandler{}

	defaultList := helpers.TodoList{Id: "default-list", Name: "Default List"}

	defaultItem := helpers.ListItem{Id: "example-item", Name: "Example Item", Completed: false, ListId: "default-list"}

	defaultList.ListItems = append(defaultList.ListItems, defaultItem)

	PrimaryListHandler.Lists = append(PrimaryListHandler.Lists, defaultList)
	PrimaryListHandler.Items = append(PrimaryListHandler.Items, defaultItem)
}
