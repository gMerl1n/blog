package jwt

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

// TokenManager provides logic for JWT & Refresh tokens generation and parsing.
type ITokenManager interface {
	NewJWT(userID string) (string, error)
	Parse(accessToken string) (string, error)
	NewRefreshToken() (string, error)
}

type TokenManager struct {
	signingKey      string
	oneDayInSeconds int
	accessTokenTTL  int
	refreshTokenTTL int
}

func NewManager(signingKey string, oneDayInSeconds, accessTokenTTL, refreshTokenTTL int) (*TokenManager, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}

	return &TokenManager{
		signingKey:      signingKey,
		oneDayInSeconds: oneDayInSeconds,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL}, nil
}

func (m *TokenManager) NewJWT(userID string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   userID,
		ExpiresAt: time.Now().Unix() + int64(m.oneDayInSeconds*m.accessTokenTTL),
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m *TokenManager) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["sub"].(string), nil
}

func (m *TokenManager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
