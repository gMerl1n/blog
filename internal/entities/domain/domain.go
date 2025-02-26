package domain

import "time"

type User struct {
	Name         string
	Email        string
	HashPassword string
}

type Post struct {
	ID        int
	Author    string
	Title     string
	Body      string
	UpdatedAt time.Time
	CreatedAt time.Time
}
