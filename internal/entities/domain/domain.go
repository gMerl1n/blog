package domain

import "time"

type User struct {
	Name         string
	Email        string
	HashPassword string
	UpdatedAt    time.Time
	CreatedAt    time.Time
}

type Post struct {
	ID        int
	Title     string
	Body      string
	UpdatedAt time.Time
	CreatedAt time.Time
}
