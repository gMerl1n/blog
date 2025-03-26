package services

import (
	"context"
	"crypto/sha1"
	"fmt"
	"strconv"

	"github.com/gMerl1n/blog/constants"
	er "github.com/gMerl1n/blog/internal/apperrors"
	"github.com/gMerl1n/blog/internal/repository"
	"github.com/gMerl1n/blog/pkg/jwt"
	"github.com/sirupsen/logrus"
)

type ServiceUser struct {
	RepoTokens   repository.IRepositoryTokens
	RepoUser     repository.IRepositoyUser
	TokenManager jwt.ITokenManager
	// accessTokenTTL  time.Duration
	// refreshTokenTTL time.Duration
	logger *logrus.Logger
}

func NewServiceUser(repoUser repository.IRepositoyUser, repoTokens repository.IRepositoryTokens, tokenManager jwt.ITokenManager, logger *logrus.Logger) *ServiceUser {
	return &ServiceUser{
		RepoTokens:   repoTokens,
		RepoUser:     repoUser,
		TokenManager: tokenManager,
		logger:       logger,
	}
}

func (s *ServiceUser) CreateUser(ctx context.Context, name, email, password, repeatPassword string) (*jwt.Tokens, error) {

	if password != repeatPassword {
		return nil, er.IncorrectRequestParams.SetCause("password and repeated password do not match")
	}

	hashPassword := s.generatePasswordHash(password)

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

func (s *ServiceUser) LoginUser(ctx context.Context, email, loginPassword string) (*jwt.Tokens, error) {

	user, err := s.RepoUser.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	isEqual := s.comparePasswords(loginPassword, user.HashPassword)
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

func (s *ServiceUser) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(constants.SaltPassword)))
}

func (s *ServiceUser) comparePasswords(loginPassword, databasePassword string) bool {

	loginPasswordHash := s.generatePasswordHash(loginPassword)
	if loginPasswordHash == databasePassword {
		return true
	} else {
		return false
	}
}

func (s *ServiceUser) RefreshTokens(ctx context.Context, refreshTokens string) (*jwt.Tokens, error) {

	userTokens, err := s.RepoTokens.GetTokens(ctx, refreshTokens)
	if err != nil {
		return nil, err
	}

	userID, err := s.TokenManager.Parse(userTokens.AccessToken)
	if err != nil {
		s.logger.Warn(fmt.Sprintf("failed to parse token from db to get userID: %s", err))
		return nil, err
	}

	userIntID, err := strconv.Atoi(userID)
	if err != nil {
		s.logger.Warn(fmt.Sprintf("failed to convert string user ID to int user ID: %s", err))
		return nil, err
	}

	newTokens, err := s.generateTokens(userIntID)
	if err != nil {
		s.logger.Warn(fmt.Sprintf("failed to generate new tokens: %s", err))
		return nil, err
	}

	if err := s.RepoTokens.SaveTokens(ctx, userIntID, newTokens); err != nil {
		s.logger.Warn(fmt.Sprintf("failed to save new tokens in db: %s", err))
		return nil, err
	}

	return newTokens, nil

}
