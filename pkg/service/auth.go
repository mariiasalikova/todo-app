package service

import (
	"github.com/mariiasalikova/todo-app/pkg/repository"
	"github.com/mariiasalikova/todo-app"
	"crypto/sha1"
	"fmt"
)

const salt = "dfyuiopoejw73hdd9o"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthServiсе(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}


func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}