package repository

import (
	quiz "github.com/ViviMemory/quiz_api"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user quiz.User) (int, error)
	GetUser(username, password string) (quiz.User, error)
}

type TodoList interface {
	Create(userId int, list quiz.TodoList) (int, error)
	GetAll(userId int) ([]quiz.TodoList, error)
	GetById(userId, listId int) (quiz.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input quiz.UpdateListInput) error
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
