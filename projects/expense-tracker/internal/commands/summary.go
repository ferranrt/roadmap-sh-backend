package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SummaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show summary of expenses",
	Long:  "Show summary of expenses",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Summary expenses")
	},
}
