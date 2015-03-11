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

func InsertUser(db *sqlx.DB, name string, email string, password string) (*User, []error) {
	u := User{}
	u.Name = name
	u.Email = email
	u.PasswordHash = password

	errors := validateUser(u)
	if len(errors) > 0 {
		return nil, errors
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, []error{err}
	}

	user, err := insertUser(db, name, email, string(hashed))
	if err != nil {
		return user, []error{err}
	} else {
		return user, nil
	}
}

//
// Validation
//

func validateUser(u User) []error {
	err := []error{}
	if len(u.Name) == 0 {
		err = append(err, errors.New("Name must not be empty"))
	}
	if len(u.Email) == 0 {
		err = append(err, errors.New("Email must not be empty"))
	}
	if len(u.PasswordHash) == 0 {
		err = append(err, errors.New("Password must not be empty"))
	}

	return err
}

//
// Data Access Methods (no real business logic or validation
//

func insertUser(db *sqlx.DB, name string, email string, password string) (*User, error) {
	user := &User{}
	err := db.QueryRowx("INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3) RETURNING *", name, email, password).StructScan(user)
	return user, dbError(err)
}
