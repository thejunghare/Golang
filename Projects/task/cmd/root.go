package cmd

import "github.com/spf13/cobra"

// Define root cmd
var RootCmd = &cobra.Command{
	Use:   "Cli task manager",
	Short: "Helps you to manage your task list",
	// This command does not have the RUN function as this does not perform any kind of action it only show task usages
}
