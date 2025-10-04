package repository

import (
	"context"
	"ferranrt/roadmap-sh/github-user-activity/internal/domain"
)

type UserActivityRepository interface {
	GetUserActivity(ctx context.Context, userID string) (domain.UserActivity, error)
}
