package service

import (
	"github.com/ellywynn/rest-api/pkg/models"
	"github.com/ellywynn/rest-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAllLists(userId int) ([]models.TodoList, error)
	GetById(userId, id int) (models.TodoList, error)
	Delete(userId, id int) error
	Update(userId int, id int, input models.UpdateListInput) error
}

type TodoItem interface {
	Create(userId int, listId int, item models.TodoItem) (int, error)
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
