package services

import (
	"context"
	"crypto/sha1"
	"fmt"

	"github.com/gMerl1n/blog/constants"
	er "github.com/gMerl1n/blog/internal/apperrors"
	"github.com/gMerl1n/blog/internal/repository"
	"github.com/sirupsen/logrus"
)

type ServiceUser struct {
	RepoUser repository.IRepositoyUser
	logger   *logrus.Logger
}

func NewServiceUser(repoUser repository.IRepositoyUser, logger *logrus.Logger) *ServiceUser {
	return &ServiceUser{
		RepoUser: repoUser,
		logger:   logger,
	}
}

func (s *ServiceUser) CreateUser(ctx context.Context, name, email, password, repeatPassword string) (int, error) {

	if password != repeatPassword {
		return 0, er.IncorrectRequestParams.SetCause("Password and repeated password do not match")
	}

	hashPassword := generatePasswordHash(password)

	userID, err := s.RepoUser.CreateUser(ctx, name, email, hashPassword)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(constants.SaltPassword)))
}

func comparePasswords(loginPassword, databasePassword string) bool {

	loginPasswordHash := generatePasswordHash(loginPassword)
	if loginPasswordHash != databasePassword {
		return true
	} else {
		return false
	}
}
