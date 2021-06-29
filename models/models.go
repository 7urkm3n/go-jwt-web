package models

import (
	"database/sql"
	"errors"

	"sample/models/users"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Users users.UserModelDB
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users: users.UserModelDB{DB: db},
	}
}
