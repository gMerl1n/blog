package services

import (
	"context"

	"github.com/gMerl1n/blog/internal/entities/domain"
	"github.com/gMerl1n/blog/internal/entities/requests"
	"github.com/gMerl1n/blog/internal/repository"
	"github.com/sirupsen/logrus"
)

type IServicePost interface {
	CreatePost(ctx context.Context, title, body string) (int, error)
	GetPostByID(ctx context.Context, postID int) (*domain.Post, error)
	GetPosts(ctx context.Context) ([]*domain.Post, error)
	UpdatePost(ctx context.Context, dataToUpdate requests.UpdatePostRequest) (int, error)
}

type IServiceUser interface {
	CreateUser(ctx context.Context, name, email, password, repeatPassword string) (int, error)
}

type Service struct {
	ServicePost IServicePost
	ServiceUser IServiceUser
}

func NewService(repo *repository.Repository, logger *logrus.Logger) *Service {
	return &Service{
		ServicePost: NewServicePost(repo.RepoPost, logger),
	}
}
