package pkg

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrInvalidToken = errors.New("invalid token")

type TokenManager struct {
	secret []byte
	ttl    time.Duration
	now    func() time.Time
}

type TokenContext struct {
	AccountID uint64
	Username  string
	ExpiresAt time.Time
}

type tokenClaims struct {
	AccountID uint64 `json:"accountId"`
	Username  string `json:"username"`
	jwt.RegisteredClaims
}

func NewTokenManager(secret string, ttl time.Duration) *TokenManager {
	return &TokenManager{secret: []byte(secret), ttl: ttl, now: time.Now}
}

func (manager *TokenManager) Sign(accountID uint64, username string) (string, time.Time, error) {
	now := manager.now().Truncate(time.Second)
	expiresAt := now.Add(manager.ttl).Truncate(time.Second)
	claims := tokenClaims{
		AccountID: accountID,
		Username:  username,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   fmt.Sprint(accountID),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(manager.secret)
	return signed, expiresAt, err
}

func (manager *TokenManager) ParseAuthorization(header string) (TokenContext, error) {
	prefix, tokenString, found := strings.Cut(strings.TrimSpace(header), " ")
	if !found || !strings.EqualFold(prefix, "Bearer") || tokenString == "" {
		return TokenContext{}, ErrInvalidToken
	}

	claims := &tokenClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (any, error) { return manager.secret, nil },
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)
	if err != nil || !token.Valid || claims.ExpiresAt == nil || claims.AccountID == 0 || claims.Username == "" {
		return TokenContext{}, ErrInvalidToken
	}
	return TokenContext{
		AccountID: claims.AccountID,
		Username:  claims.Username,
		ExpiresAt: claims.ExpiresAt.Time,
	}, nil
}
