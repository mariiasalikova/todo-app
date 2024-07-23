package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mariiasalikova/todo-app"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)

}


type TodoList interface {

}



type TodoItem interface {

}


type Repository struct {
	Authorization
	TodoItem
	TodoList

}



func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db)}
}