package auth

import (
	"fmt"
	"github.com/bwjson/fitnessApp/internal/dto"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"math/rand"
	"time"
)

type Manager interface {
	AccessTokenGen()
	RefreshTokenGen()
	//TODO: NewRefreshToken
}

type TokenManager struct {
	signingKey string
	ttl        time.Duration
}

func NewTokenManager(signingKey string, ttl time.Duration) (*TokenManager, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}
	return &TokenManager{signingKey: signingKey, ttl: ttl}, nil
}

func (m *TokenManager) AccessTokenGen(user dto.LoginInput) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(m.ttl).Unix(),
		Subject:   user.Email,
	})

	return accessToken.SignedString([]byte(m.signingKey))
}

func (m *TokenManager) RefreshTokenGen() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
