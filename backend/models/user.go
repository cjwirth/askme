package models

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id           int       `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password_hash" json:"-"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

func InsertUser(db *sqlx.DB, name string, email string, password string) (*User, error) {
	if len(name) == 0 {
		return nil, errors.New("User name must not be empty")
	}

	if len(email) == 0 {
		return nil, errors.New("User email must not be empty")
	}

	if len(password) == 0 {
		return nil, errors.New("User password must not be empty")
	}

	// TODO: what would be a good cost?
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}
	return insertUser(db, name, email, string(hashed))
}

//
// Data Access Methods (no real business logic or validation
//

func insertUser(db *sqlx.DB, name string, email string, password string) (*User, error) {
	user := &User{}
	err := db.QueryRowx("INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3) RETURNING *", name, email, password).StructScan(user)
	return user, err
}
