package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/ellywynn/rest-api/pkg/models"
	"github.com/ellywynn/rest-api/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

func (as *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return as.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	var salt string = "ajfdsij23jrh9"
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
