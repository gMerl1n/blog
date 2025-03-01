package domain

import "time"

type User struct {
	ID           int
	Name         string
	Email        string
	HashPassword string
	UpdatedAt    time.Time
	CreatedAt    time.Time
}

type Post struct {
	ID        int
	UserID    int
	Title     string
	Body      string
	UpdatedAt time.Time
	CreatedAt time.Time
}
