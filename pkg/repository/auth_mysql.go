package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/olKull/todo-app"
)

type AuthMySql struct {
	db *sqlx.DB
}

func NewAuthMySql(db *sqlx.DB) *AuthMySql {
	return &AuthMySql{db: db}
}

func (r *AuthMySql) CreateUser(user todo.User) (int, error) {

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values (\"%s\", \"%s\", \"%s\")", usersTable, user.Name, user.Username, user.Password)
	r.db.QueryRow(query)

	id, err := getLastInsertedId(r)

	if err != nil {
		return -1, err
	}

	return id, nil
}

func getLastInsertedId(r *AuthMySql) (int, error) {

	var id int

	query := fmt.Sprintf("SELECT id FROM %s ORDER BY id DESC LIMIT 1", usersTable)
	row := r.db.QueryRow(query)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthMySql) GetUser(username string, password string) (todo.User, error) {

	var user todo.User

	query := fmt.Sprintf("SELECT Id FROM %s WHERE username = \"%s\" AND password_hash = \"%s\" LIMIT 1", usersTable, username, password)

	err := r.db.Get(&user, query)

	return user, err
}