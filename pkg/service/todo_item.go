package service

import (
	"github.com/ellywynn/rest-api/pkg/models"
	"github.com/ellywynn/rest-api/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listrepo repository.TodoList) *TodoItemService {
	return &TodoItemService{
		repo:     repo,
		listRepo: listrepo,
	}
}

func (t *TodoItemService) Create(userId int, listId int, item models.TodoItem) (int, error) {
	_, err := t.listRepo.GetById(userId, listId)
	if err != nil {
		// list does not exist or does not belong to user
		return 0, err
	}

	return t.repo.Create(userId, listId, item)
}

func (t *TodoItemService) GetAllItems(userId, listId int) ([]models.TodoItem, error) {
	_, err := t.listRepo.GetById(userId, listId)
	if err != nil {
		// list does not exist or does not belong to user
		return nil, err
	}

	return t.repo.GetAllItems(userId, listId)
}

func (t *TodoItemService) GetById(userId, itemId int) (models.TodoItem, error) {
	return t.repo.GetById(userId, itemId)
}

func (t *TodoItemService) Delete(userId, itemId int) error {
	return t.repo.Delete(userId, itemId)
}

func (t *TodoItemService) Update(userId int, id int, input models.UpdateItemInput) error {
	return t.repo.Update(userId, id, input)
}
