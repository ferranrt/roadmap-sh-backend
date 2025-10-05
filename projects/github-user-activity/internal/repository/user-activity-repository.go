package repository

import (
	"context"
	"ferranrt/roadmap-sh/github-user-activity/internal/domain"
)

type UserActivityRepository interface {
	GetUserActivities(ctx context.Context, userID string, pagination domain.QueryPagination) ([]domain.UserActivity, error)
}
