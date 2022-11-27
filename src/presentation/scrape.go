package presentation

import "time"

type Article struct {
	Body       string    `json:"body"`
	CreatedAt  time.Time `json:"created_at"`
	ID         string    `json:"id"`
	LikesCount int       `json:"likes_count"`
	Title      string    `json:"title"`
	URL        string    `json:"url"`
	User       struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
}
