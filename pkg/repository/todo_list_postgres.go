package repository

import (
	"fmt"

	"github.com/ellywynn/rest-api/pkg/models"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{
		db: db,
	}
}

func (t *TodoListPostgres) Create(userId int, list models.TodoList) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	listQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(listQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	usersListsQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(usersListsQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, nil
	}

	return id, tx.Commit()
}

func (t *TodoListPostgres) GetAllLists(userId int) ([]models.TodoList, error) {
	var lists []models.TodoList

	query := fmt.Sprintf("SELECT l.id, l.title, l.description FROM %s AS l INNER JOIN %s AS ul ON l.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable)

	err := t.db.Select(&lists, query, userId)

	return lists, err
}

func (t *TodoListPostgres) GetById(userId, id int) (models.TodoList, error) {
	var list models.TodoList

	query := fmt.Sprintf("SELECT l.id, l.title, l.description FROM %s AS l "+
		"INNER JOIN %s AS ul ON l.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2",
		todoListsTable, usersListsTable)

	err := t.db.Get(&list, query, userId, id)

	return list, err
}
