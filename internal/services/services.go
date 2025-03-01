package services

import (
	"context"

	"github.com/gMerl1n/blog/internal/entities/domain"
	"github.com/gMerl1n/blog/internal/entities/requests"
	"github.com/gMerl1n/blog/internal/repository"
	"github.com/gMerl1n/blog/pkg/jwt"
	"github.com/sirupsen/logrus"
)

type IServicePost interface {
	CreatePost(ctx context.Context, title, body string, userID int) (int, error)
	GetPostByID(ctx context.Context, postID int) (*domain.Post, error)
	GetPosts(ctx context.Context) ([]*domain.Post, error)
	UpdatePost(ctx context.Context, dataToUpdate requests.UpdatePostRequest) (int, error)
}

type IServiceUser interface {
	CreateUser(ctx context.Context, name, email, password, repeatPassword string) (*jwt.Tokens, error)
	LoginUser(ctx context.Context, email, loginPassword string) (*jwt.Tokens, error)
}

type Service struct {
	ServicePost IServicePost
	ServiceUser IServiceUser
}

func NewService(repo *repository.Repository, tokenManager jwt.ITokenManager, logger *logrus.Logger) *Service {
	return &Service{
		ServicePost: NewServicePost(repo.RepoPost, logger),
		ServiceUser: NewServiceUser(repo.RepoUser, repo.RepoTokens, tokenManager, logger),
	}
}
