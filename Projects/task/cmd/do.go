package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks the task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int

		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				panic(err)
			} else {
				ids = append(ids, id)
			}
		}
		fmt.Print("Completed:",ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
