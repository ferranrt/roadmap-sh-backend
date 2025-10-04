package github

import (
	"context"
	"ferranrt/roadmap-sh/github-user-activity/internal/domain"
	"ferranrt/roadmap-sh/github-user-activity/internal/repository"
)

type GithubUserActivityRepository struct {
}

func NewGithubUserActivityRepository() repository.UserActivityRepository {
	return &GithubUserActivityRepository{}
}

func (r *GithubUserActivityRepository) GetUserActivity(ctx context.Context, userID string) (domain.UserActivity, error) {
	return domain.UserActivity{}, nil
}
