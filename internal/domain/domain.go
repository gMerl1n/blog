package domain

import "time"

type Post struct {
	ID        int
	Author    string
	Title     string
	Body      string
	UpdatedAt time.Time
	CreatedAt time.Time
}
