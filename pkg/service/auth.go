package service

import (
	"github.com/mariiasalikova/todo-app/pkg/repository"
	"github.com/mariiasalikova/todo-app"
	"crypto/sha1"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"time"
)

const (
	salt = "dfyuiopoejw73hdd9o"
	signingKey = "HiduIFY&*YWGEU$ED"
	tokenTTL = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

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

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
		ExpiresAt: time.Now().Add().Unix(),
		IssuedAt: time.Now().Unix(),
	},
	user.Id
	})

	return token.SignedString()
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}