package repository

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	er "github.com/gMerl1n/blog/internal/apperrors"
	"github.com/gMerl1n/blog/internal/entities/domain"
	"github.com/gMerl1n/blog/internal/entities/requests"
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

func (r *RepositoryPost) CreatePost(ctx context.Context, title, body string, userID int) (int, error) {

	var postID int

	query := fmt.Sprintf(
		`INSERT INTO %s (title, body, user_id)
	 	 VALUES ($1, $2)
	 	 RETURNING id`,
		postsTable,
	)

	if err := r.db.QueryRow(
		ctx,
		query,
		title,
		body,
		userID,
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

	rowsPosts, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, er.IncorrectRequestParams.SetCause(fmt.Sprintf("Cause: %s", err))
	}

	for rowsPosts.Next() {

		var post domain.Post

		if err := rowsPosts.Scan(
			&post.ID,
			&post.UserID,
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

func (r *RepositoryPost) RemovePostByID(ctx context.Context, postID int) (int, error) {

	var removedPostID int

	query := fmt.Sprintf(
		`DELETE FROM %s 
		WHERE id=$1
		RETURNING id`, postsTable,
	)

	if err := r.db.QueryRow(
		ctx,
		query,
		postID,
	).Scan(&removedPostID); err != nil {
		return 0, er.IncorrectRequest.SetCause(fmt.Sprintf("Cause: %s", err))
	}

	return removedPostID, nil

}

func (r *RepositoryPost) UpdatePostByID(ctx context.Context, data requests.UpdatePostRequest) (int, error) {

	var postUpdatedID int

	dataToUpdate := convertToSQLPatches(data)
	columnsToUpdate := strings.Join(dataToUpdate.Fields[1:], ", ")

	postToUpdateID := dataToUpdate.Args[0]
	argsToUpdate := dataToUpdate.Args[1:]

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id=%d RETURNING id`, postsTable, columnsToUpdate, postToUpdateID)

	if err := r.db.QueryRow(
		ctx,
		query,
		argsToUpdate...,
	).Scan(&postUpdatedID); err != nil {
		return 0, err
	}

	return postUpdatedID, nil
}

type SQLPatch struct {
	Fields []string
	Args   []interface{}
}

func convertToSQLPatches(resource interface{}) SQLPatch {
	var sqlPatch SQLPatch
	rType := reflect.TypeOf(resource)
	rVal := reflect.ValueOf(resource)
	numField := rType.NumField()

	sqlPatch.Fields = make([]string, 0, numField)
	sqlPatch.Args = make([]interface{}, 0, numField)

	for i := 0; i < numField; i++ {
		fType := rType.Field(i)
		fVal := rVal.Field(i)

		switch fVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if fVal.Int() != 0 {

				sqlPatch.Args = append(sqlPatch.Args, fVal.Int())

				columnsToUpdate := fmt.Sprintf(strings.ToLower(fType.Name)+"=$%d", i)

				sqlPatch.Fields = append(sqlPatch.Fields, columnsToUpdate)
			}
		case reflect.String:
			if fVal.String() != "" {
				sqlPatch.Args = append(sqlPatch.Args, fVal.String())

				snakeCaseColumn := toSnakeCase(fType.Name)

				columnsToUpdate := fmt.Sprintf(snakeCaseColumn+"=$%d", i)

				sqlPatch.Fields = append(sqlPatch.Fields, columnsToUpdate)
			}

		}

	}

	return sqlPatch
}

func toSnakeCase(str string) string {

	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
