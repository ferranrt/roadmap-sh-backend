package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all expenses",
	Long:  "List all expenses",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List expenses")
	},
}
