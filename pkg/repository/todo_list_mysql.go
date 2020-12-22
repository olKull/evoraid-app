package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/olKull/todo-app"
)

type TodoListMySql struct {
	db *sqlx.DB
}

func NewTodoListMySql(db *sqlx.DB) *TodoListMySql {
	return &TodoListMySql{db: db}
}

func (r *TodoListMySql) Create(userId int, list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var listId int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ('%s', '%s')", todoListsTable, list.Title, list.Description)
	r.db.QueryRow(createListQuery)

	id, err := getLastInsertedListId(r)

	if  err != nil {
		tx.Rollback()
		return 0, nil
	}

	listId = id

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ('%d', '%d')", usersListsTable, userId, listId)
	_, err = tx.Exec(createUsersListQuery)
	if err != nil {
		tx.Rollback()
		return 0, nil
	}

	return id, tx.Commit()
}

func getLastInsertedListId(r *TodoListMySql) (int, error) {

	var id int

	query := fmt.Sprintf("SELECT id FROM %s ORDER BY id DESC LIMIT 1", todoListsTable)
	row := r.db.QueryRow(query)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TodoListMySql) GetAll(userId int) ([]todo.TodoList, error) {

	var lists []todo.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s as tl INNER JOIN %s as ul ON tl.id = ul.list_id WHERE ul.user_id = %d", todoListsTable, usersListsTable, userId)
	err := r.db.Select(&lists, query)

	return lists, err

}

func (r *TodoListMySql) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description 
 								 FROM %s as tl INNER JOIN %s as ul ON tl.id = ul.list_id 
								 WHERE ul.user_id = %d AND ul.list_id = %d`, todoListsTable, usersListsTable, userId, listId)
	err := r.db.Get(&list, query)

	return list, err

}