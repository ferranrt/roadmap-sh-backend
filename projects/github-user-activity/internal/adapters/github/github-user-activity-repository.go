package github

import (
	"context"
	"encoding/json"
	"ferranrt/roadmap-sh/github-user-activity/internal/domain"
	"ferranrt/roadmap-sh/github-user-activity/internal/repository"
	"fmt"
	"io"
	"net/http"
)

func generateGithubActivityUrl(userID string) string {
	return fmt.Sprintf("https://api.github.com/users/%s/events/public", userID)
}

type GithubUserActivityResponse []GithubUserActivityEvent

type GithubUserActivityEvent struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Public    bool   `json:"public"`
	CreatedAt string `json:"created_at"`
	Actor     struct {
		Login        string `json:"login"`
		ID           int    `json:"id"`
		AvatarURL    string `json:"avatar_url"`
		URL          string `json:"url"`
		GravatarID   string `json:"gravatar_id"`
		DisplayLogin string `json:"display_login"`
	} `json:"actor"`
	Repo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"repo"`
	Payload struct {
		Action string `json:"action"`
	} `json:"payload"`
}

type GithubUserActivityRepository struct {
}

func NewGithubUserActivityRepository() repository.UserActivityRepository {
	return &GithubUserActivityRepository{}
}

func (r *GithubUserActivityRepository) GetUserActivity(ctx context.Context, userID string) (domain.UserActivity, error) {
	url := generateGithubActivityUrl(userID)
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		return domain.UserActivity{}, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return domain.UserActivity{}, err
	}
	var githubUserActivityResponse GithubUserActivityResponse
	err = json.Unmarshal(body, &githubUserActivityResponse)
	if err != nil {
		return domain.UserActivity{}, err
	}
	for _, event := range githubUserActivityResponse {
		fmt.Println("--------------------------------")
		fmt.Println(event)
		fmt.Println("--------------------------------")

	}
	return domain.UserActivity{}, nil
}
