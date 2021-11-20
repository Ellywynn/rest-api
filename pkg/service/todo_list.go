package service

import (
	"github.com/ellywynn/rest-api/pkg/models"
	"github.com/ellywynn/rest-api/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{
		repo: repo,
	}
}

func (t *TodoListService) Create(userId int, list models.TodoList) (int, error) {
	return t.repo.Create(userId, list)
}

func (t *TodoListService) GetAllLists(userId int) ([]models.TodoList, error) {
	return t.repo.GetAllLists(userId)
}

func (t *TodoListService) GetById(userId, id int) (models.TodoList, error) {
	return t.repo.GetById(userId, id)
}

func (t *TodoListService) Delete(userId, id int) error {
	return t.repo.Delete(userId, id)
}

func (t *TodoListService) Update(userId int, id int, input models.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return t.repo.Update(userId, id, input)
}
