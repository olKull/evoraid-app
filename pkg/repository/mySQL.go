package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
	todoListsTable = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable = "todo_items"
	listsItemTable = "lists_items"
)

type Config struct {
	Host string
	Port string
	Username string
	Password string
	DBName string
}

func NewMySqlDb(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}