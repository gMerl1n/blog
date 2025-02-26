package repository

import (
	"fmt"

	er "github.com/gMerl1n/blog/internal/apperrors"
	"github.com/gMerl1n/blog/pkg/jwt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

const (
	tokensTable = "tokens"
)

type RepositoryTokens struct {
	db     *pgxpool.Pool
	logger *logrus.Logger
}

func NewRepositoryTokens(db *pgxpool.Pool, logger *logrus.Logger) *RepositoryTokens {
	return &RepositoryTokens{
		db:     db,
		logger: logger,
	}
}

func (r *RepositoryTokens) SaveTokens(ctx context.Context, userID int, tokens *jwt.Tokens) error {

	query := fmt.Sprintf(
		`INSERT INTO %s (token, refresh_token)
	 	 VALUES ($1, $2)
	 	 RETURNING id`,
		tokensTable,
	)

	if err := r.db.QueryRow(
		ctx,
		query,
		tokens.AccessToken,
		tokens.RefreshToken,
	); err != nil {
		return er.IncorrectRequestParams.SetCause(fmt.Sprintf("Cause: %s", err))

	}

	return nil

}

func (r *RepositoryTokens) GetTokens(ctx context.Context, refreshToken string) (*jwt.Tokens, error) {

	var tokens jwt.Tokens

	query := fmt.Sprintf(
		`SELECT token, refresh_token
		FROM %s
		WHERE refresh_token = $1`,
		tokensTable,
	)

	if err := r.db.QueryRow(
		ctx,
		query,
		refreshToken,
	).Scan(&tokens.AccessToken, &tokens.RefreshToken); err != nil {
		return nil, er.IncorrectRequestParams.SetCause(fmt.Sprintf("Probably, no such a refresh token in db: %s", refreshToken))
	}

	return &tokens, nil

}
