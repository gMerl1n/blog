package repository

import (
	"context"

	"github.com/gMerl1n/blog/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IRepositoryPost interface {
	CreatePost(ctx context.Context, title, body string) (int, error)
	GetPostByID(ctx context.Context, postID int) (*domain.Post, error)
}

type Repository struct {
	RepoPost IRepositoryPost
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		RepoPost: NewRepositoryPost(db),
	}
}
