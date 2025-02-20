package services

import (
	"github.com/gMerl1n/blog/internal/domain"
	"github.com/gMerl1n/blog/internal/repository"
)

type IServicePost interface {
	CreatePost(title, body string) (int, error)
	GetPostByID(postID int) (*domain.Post, error)
}

type Service struct {
	ServicePost IServicePost
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		ServicePost: NewServicePost(repo.RepoPost),
	}
}
