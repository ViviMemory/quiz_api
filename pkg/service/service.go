package service

import (
	quiz "github.com/ViviMemory/quiz_api"
	"github.com/ViviMemory/quiz_api/pkg/repository"
)

type Authorization interface {
	CreateUser(user quiz.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
