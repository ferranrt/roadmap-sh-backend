package commands

import (
	"context"
	"ferranrt/roadmap-sh/github-user-activity/internal/adapters/github"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	user_id string
)

var rootCmd = &cobra.Command{
	Use:   "github-user-activity",
	Short: "GitHub User Activity",
	Long:  "GitHub User Activity",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GitHub User Activity")
		if user_id == "" {
			fmt.Println("User ID is required")
			return
		}
		activityRepository := github.NewGithubUserActivityRepository()
		activity, err := activityRepository.GetUserActivity(context.Background(), user_id)
		if err != nil {
			fmt.Println("Error getting user activity:", err)
			return
		}
		fmt.Println(activity)
	},
}

func setupRootCmd(c *cobra.Command) {
	c.Flags().StringVarP(&user_id, "user", "u", "", "")
}

func GetRootCmd() *cobra.Command {
	setupRootCmd(rootCmd)
	return rootCmd
}
