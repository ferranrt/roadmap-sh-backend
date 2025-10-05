package github

import (
	"context"
	"encoding/json"
	"ferranrt/roadmap-sh/github-user-activity/internal/domain"
	"ferranrt/roadmap-sh/github-user-activity/internal/repository"
	"ferranrt/roadmap-sh/github-user-activity/internal/utils"
	"fmt"
	"io"
	"net/http"
)

func generateGithubActivityUrl(userID string, limit int, page int) string {
	return fmt.Sprintf("https://api.github.com/users/%s/events/public?per_page=%d&page=%d", userID, limit, page)
}

type GithubUserActivityRepository struct{}

func NewGithubUserActivityRepository() repository.UserActivityRepository {
	return &GithubUserActivityRepository{}
}

func dtoToDomain(event GithubUserActivityEvent) domain.UserActivity {
	var output domain.UserActivity
	output.Type = event.Type
	output.Actor.User = event.Actor.Login
	output.Repo.Name = event.Repo.Name
	return output
}

func (r *GithubUserActivityRepository) GetUserActivities(ctx context.Context, userID string, pagination domain.QueryPagination) ([]domain.UserActivity, error) {
	safeLimit := utils.Min(pagination.Limit, MAX_EVENTS_PER_PAGE)
	url := generateGithubActivityUrl(userID, safeLimit, pagination.Page)
	response, err := http.Get(url)
	if err != nil {
		return []domain.UserActivity{}, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []domain.UserActivity{}, err
	}
	var githubUserActivityResponse GithubUserActivityResponse
	err = json.Unmarshal(body, &githubUserActivityResponse)
	if err != nil {
		return []domain.UserActivity{}, err
	}
	userActivities := make([]domain.UserActivity, len(githubUserActivityResponse))
	for _, event := range githubUserActivityResponse {
		userActivities = append(userActivities, dtoToDomain(event))

	}
	return userActivities, nil
}
