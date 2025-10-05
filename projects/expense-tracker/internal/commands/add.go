package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new expense",
	Long:  "Add a new expense",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add expense")
	},
}
