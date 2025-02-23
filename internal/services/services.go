package services

import (
	"context"

	"github.com/gMerl1n/blog/internal/domain"
	"github.com/gMerl1n/blog/internal/repository"
	"github.com/sirupsen/logrus"
)

type IServicePost interface {
	CreatePost(ctx context.Context, title, body string) (int, error)
	GetPostByID(ctx context.Context, postID int) (*domain.Post, error)
	GetPosts(ctx context.Context) ([]*domain.Post, error)
}

type Service struct {
	ServicePost IServicePost
}

func NewService(repo *repository.Repository, logger *logrus.Logger) *Service {
	return &Service{
		ServicePost: NewServicePost(repo.RepoPost, logger),
	}
}
