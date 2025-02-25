package repository

import (
	"context"

	"github.com/gMerl1n/blog/internal/entities/domain"
	"github.com/gMerl1n/blog/internal/entities/requests"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type IRepositoryPost interface {
	CreatePost(ctx context.Context, title, body string) (int, error)
	GetPostByID(ctx context.Context, postID int) (*domain.Post, error)
	GetPosts(ctx context.Context) ([]*domain.Post, error)
	UpdatePostByID(ctx context.Context, data requests.UpdatePostRequest) (int, error)
}

type Repository struct {
	RepoPost IRepositoryPost
}

func NewRepository(db *pgxpool.Pool, logger *logrus.Logger) *Repository {
	return &Repository{
		RepoPost: NewRepositoryPost(db, logger),
	}
}
