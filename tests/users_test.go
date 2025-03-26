package tests

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/gMerl1n/blog/internal/entities/requests"
	"github.com/gavv/httpexpect/v2"
)

const (
	host = "localhost:9004"
)

func TestCreateUser(t *testing.T) {

	// Arrange

	u := url.URL{
		Scheme: "http",
		Host:   host,
	}

	e := httpexpect.Default(t, u.String())

	// Act
	// Assert
	e.POST("/api/users/").
		WithJSON(requests.CreateUserRequest{
			Name:           "John2",
			Email:          "John2@John123.com",
			Password:       "123456789",
			RepeatPassword: "123456789",
		}).
		Expect().
		Status(http.StatusCreated).
		JSON().Object().
		ContainsKey("AccessToken").
		ContainsKey("RefreshToken")

}

func TestCreateUserWrongRepeatPassword(t *testing.T) {

	u := url.URL{
		Scheme: "http",
		Host:   host,
	}

	e := httpexpect.Default(t, u.String())

	// Act
	// Assert
	e.POST("/api/users/").
		WithJSON(requests.CreateUserRequest{
			Name:           "John2",
			Email:          "John2@John123.com",
			Password:       "123456789",
			RepeatPassword: "1",
		}).
		Expect().
		Status(400)

}

func TestCreateUserIncorrectInput(t *testing.T) {

	u := url.URL{
		Scheme: "http",
		Host:   host,
	}

	e := httpexpect.Default(t, u.String())

	// Act
	// Assert
	e.POST("/api/users/").
		WithJSON(requests.CreateUserRequest{
			Name:           "John2",
			Email:          "John2@John123.com",
			Password:       "123456789",
			RepeatPassword: "123456789",
		}).
		Expect().
		Status(http.StatusConflict)

}

func TestLoginUser(t *testing.T) {

	// Arrange

	u := url.URL{
		Scheme: "http",
		Host:   host,
	}

	e := httpexpect.Default(t, u.String())

	// Act
	// Assert
	e.POST("/api/users/login").
		WithJSON(requests.LoginUserRequest{
			Email:    "John2@John123.com",
			Password: "123456789",
		}).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ContainsKey("AccessToken").
		ContainsKey("RefreshToken")
}
