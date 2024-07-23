package service

import (
	"github.com/mariiasalikova/todo-app/pkg/repository"
	"github.com/mariiasalikova/todo-app"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
}


type TodoList interface {

}



type TodoItem interface {

}


type Service struct {
	Authorization
	TodoItem
	TodoList

}



func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthServiсе(repos.Authorization)}
}