package github

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
