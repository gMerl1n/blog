package services

import (
	"github.com/gMerl1n/blog/internal/domain"
	"github.com/gMerl1n/blog/internal/repository"
)

type ServicePost struct {
	repoPost repository.IRepositoryPost
}

func NewServicePost(repoPost repository.IRepositoryPost) *ServicePost {
	return &ServicePost{repoPost: repoPost}
}

func (s *ServicePost) CreatePost(title, body string) (int, error) {
	postID, err := s.repoPost.CreatePost(title, body)
	if err != nil {
		return 0, err
	}

	return postID, err
}

func (s *ServicePost) GetPostByID(postID int) (*domain.Post, error) {
	post, err := s.repoPost.GetPostByID(postID)
	if err != nil {
		return nil, err
	}

	return post, err
}
