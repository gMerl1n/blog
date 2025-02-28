package services

import (
	"context"
	"crypto/sha1"
	"fmt"
	"strconv"
	"time"

	"github.com/gMerl1n/blog/constants"
	er "github.com/gMerl1n/blog/internal/apperrors"
	"github.com/gMerl1n/blog/internal/repository"
	"github.com/gMerl1n/blog/pkg/jwt"
	"github.com/sirupsen/logrus"
)

type ServiceUser struct {
	RepoTokens      repository.IRepositoryTokens
	RepoUser        repository.IRepositoyUser
	TokenManager    jwt.ITokenManager
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	logger          *logrus.Logger
}

func NewServiceUser(repoUser repository.IRepositoyUser, repoTokens repository.IRepositoryTokens, logger *logrus.Logger) *ServiceUser {
	return &ServiceUser{
		RepoTokens: repoTokens,
		RepoUser:   repoUser,
		logger:     logger,
	}
}

func (s *ServiceUser) CreateUser(ctx context.Context, name, email, password, repeatPassword string) (*jwt.Tokens, error) {

	if password != repeatPassword {
		return nil, er.IncorrectRequestParams.SetCause("Password and repeated password do not match")
	}

	hashPassword := generatePasswordHash(password)

	userID, err := s.RepoUser.CreateUser(ctx, name, email, hashPassword)
	if err != nil {
		return nil, err
	}

	tokens, err := s.generateTokens(userID)
	if err != nil {
		return nil, err
	}

	if err := s.RepoTokens.SaveTokens(ctx, userID, tokens); err != nil {
		return nil, err
	}

	return tokens, nil
}

func (s *ServiceUser) LoginUser(ctx context.Context, email, password string) (*jwt.Tokens, error) {

	user, err := s.RepoUser.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	isEqual := comparePasswords(user.HashPassword, password)
	if !isEqual {
		return nil, fmt.Errorf("passwords do not match")
	}

	tokens, err := s.generateTokens(user.ID)
	if err != nil {
		return nil, err
	}

	if err := s.RepoTokens.SaveTokens(ctx, user.ID, tokens); err != nil {
		return nil, err
	}

	return tokens, nil

}

func (s *ServiceUser) generateTokens(userID int) (*jwt.Tokens, error) {
	var (
		tokens jwt.Tokens
		err    error
	)

	tokens.AccessToken, err = s.TokenManager.NewJWT(strconv.Itoa(userID))
	if err != nil {
		return nil, err
	}

	tokens.RefreshToken, err = s.TokenManager.NewRefreshToken()
	if err != nil {
		return nil, err
	}

	return &tokens, err
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
