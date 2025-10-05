package domain

type UserActivity struct {
	Type  string `json:"type"`
	Actor struct {
		User string `json:"user"`
	}
	Repo struct {
		Name string `json:"name"`
	}
}

type UserActivityPrinter interface {
	Print(userActivity UserActivity)
}
