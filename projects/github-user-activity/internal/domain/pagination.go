package domain

type QueryPagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"` // Offset
}
