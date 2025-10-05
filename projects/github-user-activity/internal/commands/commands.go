package commands

import (
	"context"
	"ferranrt/roadmap-sh/github-user-activity/internal/adapters/github"
	"ferranrt/roadmap-sh/github-user-activity/internal/domain"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	user_id  string
	offset   int
	pageSize int
)

var rootCmd = &cobra.Command{
	Use:   "github-user-activity",
	Short: "GitHub User Activity",
	Long:  "GitHub User Activity",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GitHub User Activity")
		if offset == 0 {
			offset = github.MAX_EVENTS_PER_PAGE
		}
		if pageSize < 0 {
			pageSize = 0
		}
		if user_id == "" {
			fmt.Println("User ID is required")
			return
		}
		activityRepository := github.NewGithubUserActivityRepository()
		pagination := domain.QueryPagination{
			Limit: offset,
			Page:  pageSize,
		}
		activities, err := activityRepository.GetUserActivities(context.Background(), user_id, pagination)
		if err != nil {
			fmt.Println("Error getting user activity:", err)
			return
		}
		fmt.Println("Output:")
		for _, activity := range activities {
			printer := github.GetActivityPrinter(activity)
			printer.Print(activity)
		}
	},
}

func setupRootCmd(c *cobra.Command) {
	c.Flags().StringVarP(&user_id, "user", "u", "", "")
	c.Flags().IntVarP(&offset, "offset", "o", 0, "")
	c.Flags().IntVarP(&pageSize, "page", "p", 0, "")
}

func GetRootCmd() *cobra.Command {
	setupRootCmd(rootCmd)
	return rootCmd
}
