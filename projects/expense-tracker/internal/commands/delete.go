package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an expense",
	Long:  "Delete an expense",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete expense")
	},
}
