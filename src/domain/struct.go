package domain

import "time"

type Article struct {
	Body       string
	CreatedAt  time.Time
	ID         string
	LikesCount int
	Title      string
	URL        string
	User
}

type User struct {
	ID   string
	Name string
}
