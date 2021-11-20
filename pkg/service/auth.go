package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ellywynn/rest-api/pkg/models"
	"github.com/ellywynn/rest-api/pkg/repository"
)

const (
	tokenExpire = 12 * time.Hour
)

var (
	tokenKey = []byte(os.Getenv("TOKEN_KEY"))
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

func (as *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return as.repo.CreateUser(user)
}

func (as *AuthService) GenerateToken(email, password string) (string, error) {
	// Get user from DB
	user, err := as.repo.GetUser(email, generatePasswordHash(password))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenExpire).Unix(),
		IssuedAt:  time.Now().Unix(),
	},
		user.Id,
	})

	return token.SignedString(tokenKey)
}

func (as *AuthService) ParseToken(token string) (int, error) {
	t, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid sign method")
		}

		return tokenKey, nil
	})

	if err != nil {
		return 0, nil
	}

	claims, ok := t.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	var salt string = "ajfdsij23jrh9"
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
