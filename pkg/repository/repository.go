package repository

import (
	"github.com/ellywynn/rest-api/pkg/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(email, password string) (models.User, error)
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
	GetAllItems(userId int, listId int) ([]models.TodoItem, error)
	GetById(userId, id int) (models.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId int, id int, input models.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
