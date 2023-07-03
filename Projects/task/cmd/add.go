package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds the task to task list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("add called")
	},
}

/* init() is a function that runs before the main function */
func init() {
	RootCmd.AddCommand(addCmd)
}