package repository

import (
	"context"
	"fmt"
	"strings"

	er "github.com/gMerl1n/blog/internal/apperrors"
	"github.com/gMerl1n/blog/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

const (
	postsTable = "posts"
)

type RepositoryPost struct {
	db     *pgxpool.Pool
	logger *logrus.Logger
}

func NewRepositoryPost(db *pgxpool.Pool, logger *logrus.Logger) *RepositoryPost {
	return &RepositoryPost{
		db:     db,
		logger: logger,
	}
}

func (r *RepositoryPost) CreatePost(ctx context.Context, title, body string) (int, error) {

	var postID int

	query := fmt.Sprintf(
		`INSERT INTO %s (title, body)
	 	 VALUES ($1, $2)
	 	 RETURNING id`,
		postsTable,
	)

	if err := r.db.QueryRow(
		ctx,
		query,
		title,
		body,
	).Scan(&postID); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return 0, er.PostIsAlready.SetCause(fmt.Sprintf("Cause: %s", err))
		} else {
			return 0, er.IncorrectRequestParams.SetCause(fmt.Sprintf("Cause: %s", err))
		}

	}

	return postID, nil

}

func (r *RepositoryPost) GetPostByID(ctx context.Context, postID int) (*domain.Post, error) {

	var post domain.Post

	query := `SELECT *
		 	  FROM posts
		 	  WHERE id = $1`

	if err := r.db.QueryRow(
		ctx,
		query,
		postID,
	).Scan(&post.ID, &post.Title, &post.Body, &post.UpdatedAt, &post.CreatedAt); err != nil {
		return nil, er.IncorrectRequest.SetCause(fmt.Sprintf("Cause: %s", err))
	}

	return &post, nil
}

func (r *RepositoryPost) GetPosts(ctx context.Context) ([]*domain.Post, error) {

	listPosts := make([]*domain.Post, 0)

	query := `SELECT * 
			  FROM posts`

	rowsCargos, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, er.IncorrectRequestParams.SetCause(fmt.Sprintf("Cause: %s", err))
	}

	for rowsCargos.Next() {

		var post domain.Post

		if err := rowsCargos.Scan(
			&post.ID,
			&post.Title,
			&post.Body,
			&post.UpdatedAt,
			&post.CreatedAt,
		); err != nil {
			return nil, er.IncorrectData.SetCause(fmt.Sprintf("Cause: %s", err))
		}

		listPosts = append(listPosts, &post)

	}

	return listPosts, nil

}
