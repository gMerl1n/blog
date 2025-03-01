package repository

import (
	"context"
	"fmt"
	"strings"

	er "github.com/gMerl1n/blog/internal/apperrors"
	"github.com/gMerl1n/blog/internal/entities/domain"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

const (
	usersTable = "users"
)

type RepositoryUser struct {
	db     *pgxpool.Pool
	logger *logrus.Logger
}

func NewRepositoryUser(db *pgxpool.Pool, logger *logrus.Logger) *RepositoryUser {
	return &RepositoryUser{
		db:     db,
		logger: logger,
	}
}

func (r *RepositoryUser) CreateUser(ctx context.Context, name, email, hashPassword string) (int, error) {

	var userID int

	query := fmt.Sprintf(
		`INSERT INTO %s (name, email, hash_password)
	 	VALUES ($1, $2, $3)
	 	RETURNING id`,
		usersTable,
	)

	if err := r.db.QueryRow(
		ctx,
		query,
		name,
		email,
		hashPassword,
	).Scan(&userID); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return 0, er.PostIsAlready.SetCause(fmt.Sprintf("Cause: %s", err))
		} else {
			return 0, er.IncorrectRequestParams.SetCause(fmt.Sprintf("Cause: %s", err))
		}

	}

	return userID, nil

}

func (r *RepositoryUser) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {

	var user domain.User

	query := fmt.Sprintf(
		`SELECT * 
		FROM %s 
		WHERE email=$1`, usersTable,
	)

	if err := r.db.QueryRow(
		ctx,
		query,
		email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.HashPassword, &user.UpdatedAt, &user.CreatedAt); err != nil {
		return nil, er.IncorrectRequest.SetCause(fmt.Sprintf("Cause: %s", err))
	}

	return &user, nil

}
