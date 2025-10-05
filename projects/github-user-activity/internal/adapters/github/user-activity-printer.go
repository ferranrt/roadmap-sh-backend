package github

import (
	"ferranrt/roadmap-sh/github-user-activity/internal/domain"
	"fmt"
)

type DefaultUserActivityPrinter struct{}

func (p *DefaultUserActivityPrinter) Print(activity domain.UserActivity) {
	fmt.Println("A", activity)
}

func NewDefaultUserActivityPrinter() domain.UserActivityPrinter {
	return &DefaultUserActivityPrinter{}
}

type PushEventPrinter struct{}

func (p *PushEventPrinter) Print(activity domain.UserActivity) {
	fmt.Printf("- %s Pushed to %s  \n", activity.Actor.User, activity.Repo.Name)
}

func NewPushEventPrinter() domain.UserActivityPrinter {
	return &PushEventPrinter{}
}

type PullRequestEventPrinter struct{}

func (p *PullRequestEventPrinter) Print(activity domain.UserActivity) {
	fmt.Printf("- %s Pulled request to %s  \n", activity.Actor.User, activity.Repo.Name)
}

func NewPullRequestEventPrinter() domain.UserActivityPrinter {
	return &PullRequestEventPrinter{}
}

type WatchEventPrinter struct{}

func (p *WatchEventPrinter) Print(activity domain.UserActivity) {
	fmt.Printf("- %s Watched %s  \n", activity.Actor.User, activity.Repo.Name)
}

func NewWatchEventPrinter() domain.UserActivityPrinter {
	return &WatchEventPrinter{}
}

type IssueCommentEventPrinter struct{}

func (p *IssueCommentEventPrinter) Print(activity domain.UserActivity) {
	fmt.Printf("- %s Commented on %s  \n", activity.Actor.User, activity.Repo.Name)
}

func NewIssueCommentEventPrinter() domain.UserActivityPrinter {
	return &IssueCommentEventPrinter{}
}

type IssueEventPrinter struct{}

func (p *IssueEventPrinter) Print(activity domain.UserActivity) {
	fmt.Printf("- %s Opened an issue on %s  \n", activity.Actor.User, activity.Repo.Name)
}

func NewIssueEventPrinter() domain.UserActivityPrinter {
	return &IssueEventPrinter{}
}

type ForkEventPrinter struct{}

func (p *ForkEventPrinter) Print(activity domain.UserActivity) {
	fmt.Printf("- %s Forked %s  \n", activity.Actor.User, activity.Repo.Name)
}

func NewForkEventPrinter() domain.UserActivityPrinter {
	return &ForkEventPrinter{}
}

type PullRequestReviewCommentEventPrinter struct{}

func (p *PullRequestReviewCommentEventPrinter) Print(activity domain.UserActivity) {
	fmt.Printf("- %s Reviewed a pull request on %s  \n", activity.Actor.User, activity.Repo.Name)
}

func NewPullRequestReviewCommentEventPrinter() domain.UserActivityPrinter {
	return &PullRequestReviewCommentEventPrinter{}
}

type PullRequestReviewEventPrinter struct{}

func (p *PullRequestReviewEventPrinter) Print(activity domain.UserActivity) {
	fmt.Printf("- %s Reviewed a pull request on %s  \n", activity.Actor.User, activity.Repo.Name)
}

func NewPullRequestReviewEventPrinter() domain.UserActivityPrinter {
	return &PullRequestReviewEventPrinter{}
}

func GetActivityPrinter(activity domain.UserActivity) domain.UserActivityPrinter {
	if activity.Type == "PushEvent" {
		return NewPushEventPrinter()
	}
	if activity.Type == "PullRequestEvent" {
		return NewPullRequestEventPrinter()
	}
	if activity.Type == "WatchEvent" {
		return NewWatchEventPrinter()
	}
	if activity.Type == "IssueCommentEvent" {
		return NewIssueCommentEventPrinter()
	}
	if activity.Type == "IssuesEvent" {
		return NewIssueEventPrinter()
	}
	if activity.Type == "ForkEvent" {
		return NewForkEventPrinter()
	}
	if activity.Type == "PullRequestReviewCommentEvent" {
		return NewPullRequestReviewCommentEventPrinter()
	}
	if activity.Type == "PullRequestReviewEvent" {
		return NewPullRequestReviewEventPrinter()
	}
	return NewDefaultUserActivityPrinter()
}
