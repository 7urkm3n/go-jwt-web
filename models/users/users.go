package users

import (
	"database/sql"
	"time"
)

type UserModelDB struct {
	DB *sql.DB
}

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
}

var AnonymousUser = &User{}

func (m UserModelDB) GetByEmail(email string) (*User, error) {
	var user User
	err := m.DB.QueryRow(`SELECT id, email, created_at FROM users WHERE email == $1`, email).
		Scan(&user.ID, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m UserModelDB) GetByEmailWithPassword(email string) (*User, error) {
	var user User
	err := m.DB.QueryRow(`SELECT * FROM users WHERE email == $1`, email).
		Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m UserModelDB) Insert(user *User) (*User, error) {
	return nil, nil
}

func (u *User) IsAnonymous() bool {
	return u == AnonymousUser
}
