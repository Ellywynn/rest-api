package repository

import (
	"fmt"
	"strings"

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

	query := fmt.Sprintf("SELECT l.* FROM %s AS l INNER JOIN %s AS ul ON l.id = ul.list_id WHERE ul.user_id = $1",
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

func (t *TodoListPostgres) Delete(userId, id int) error {
	query := fmt.Sprintf("DELETE FROM %s l USING %s ul WHERE l.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2",
		todoListsTable, usersListsTable)

	_, err := t.db.Exec(query, userId, id)

	return err
}

func (t *TodoListPostgres) Update(userId int, id int, input models.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("description=%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s l SET %s FROM %s ul WHERE l.id = ul.list_id AND ul.list_id=%d AND ul.user_id=%d",
		todoListsTable, setQuery, usersListsTable, argId, argId+1)

	args = append(args, id, userId)

	_, err := t.db.Exec(query, args...)
	return err
}
