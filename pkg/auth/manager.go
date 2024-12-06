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
	//TODO: ParseToken()
	//TODO: NewRefreshToken()
}

type TokenManager struct {
	signingKey string
}

func NewTokenManager(signingKey string) (*TokenManager, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}
	return &TokenManager{signingKey: signingKey}, nil
}

func (m *TokenManager) AccessTokenGen(user dto.LoginInput, accessTokenTTL time.Duration) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(accessTokenTTL).Unix(),
		Subject:   user.Email,
	})

	return accessToken.SignedString([]byte(m.signingKey))
}

func (m *TokenManager) RefreshTokenGen(refreshTokenTTL time.Duration) (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
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
