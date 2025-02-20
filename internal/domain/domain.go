package domain

import "time"

type Post struct {
	ID        int
	UserID    int
	Title     string
	Body      string
	UpdatedAt time.Time
	CreatedAt time.Time
}
