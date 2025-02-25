package services

import (
	"context"

	"github.com/gMerl1n/blog/internal/entities/domain"
	"github.com/gMerl1n/blog/internal/entities/requests"
	"github.com/gMerl1n/blog/internal/repository"
	"github.com/sirupsen/logrus"
)

type ServicePost struct {
	repoPost repository.IRepositoryPost
	logger   *logrus.Logger
}

func NewServicePost(repoPost repository.IRepositoryPost, logger *logrus.Logger) *ServicePost {
	return &ServicePost{
		repoPost: repoPost,
		logger:   logger,
	}
}

func (s *ServicePost) CreatePost(ctx context.Context, title, body string) (int, error) {
	postID, err := s.repoPost.CreatePost(ctx, title, body)
	if err != nil {
		return 0, err
	}

	return postID, err
}

func (s *ServicePost) GetPostByID(ctx context.Context, postID int) (*domain.Post, error) {
	post, err := s.repoPost.GetPostByID(ctx, postID)
	if err != nil {
		return nil, err
	}

	return post, err
}

func (s *ServicePost) GetPosts(ctx context.Context) ([]*domain.Post, error) {
	listPosts, err := s.repoPost.GetPosts(ctx)
	if err != nil {
		return nil, err
	}

	return listPosts, nil
}

func (s *ServicePost) UpdatePost(ctx context.Context, dataToUpdate requests.UpdatePostRequest) (int, error) {
	postUpdatedID, err := s.repoPost.UpdatePostByID(ctx, dataToUpdate)
	if err != nil {
		return 0, err
	}

	return postUpdatedID, nil
}
