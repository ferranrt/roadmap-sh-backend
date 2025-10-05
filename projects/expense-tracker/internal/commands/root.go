package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "expense-tracker",
	Short: "Expense tracker is a CLI tool for tracking your expenses",
	Long:  "Expense tracker is a CLI tool for tracking your expenses",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
}

func injectCommands(cmd *cobra.Command) {
	cmd.AddCommand(AddCmd)
	cmd.AddCommand(ListCmd)
	cmd.AddCommand(SummaryCmd)
	cmd.AddCommand(DeleteCmd)
}

func GetRootCmd() *cobra.Command {
	injectCommands(RootCmd)

	return RootCmd
}
