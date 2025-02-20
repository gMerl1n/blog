package repository

import (
	"github.com/gMerl1n/blog/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryPost struct {
	db *pgxpool.Pool
}

func NewRepositoryPost(db *pgxpool.Pool) *RepositoryPost {
	return &RepositoryPost{db: db}
}

func (r *RepositoryPost) CreatePost(title, body string) (int, error) {
	return 0, nil
}

func (r *RepositoryPost) GetPostByID(postID int) (*domain.Post, error) {
	return nil, nil
}
