package models

import (
	"database/sql"
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

func (u *User) CheckPassword(password string) bool {
	return passwordIsCorrect(u.PasswordHash, password)
}

func GetUserById(db *sqlx.DB, id int) *User {
	user, _ := getUserById(db, id)
	return user
}

func GetUserByName(db *sqlx.DB, name string) *User {
	user, _ := getUserByName(db, name)
	return user
}

// InsertUser inserts a user into the database.
// Returns errors that occur -- validation and data errors
func InsertUser(db *sqlx.DB, name string, email string, password string) (*User, error) {
	u := User{}
	u.Name = name
	u.Email = email
	u.PasswordHash = password

	hashed, vErr := hashPassword(password)
	if vErr != nil {
		err := NewValidationError()
		err.AddReason("Password has bad format")
		return nil, err
	}

	if err := ValidateUser(u); err != nil {
		return nil, err
	}

	user, err := insertUser(db, name, email, hashed)
	return user, err

}

//
// Validation
//

func ValidateUser(u User) error {
	err := NewValidationError()
	if len(u.Name) == 0 {
		err.AddReason("Name must not be empty.")
	}
	if len(u.Email) == 0 {
		err.AddReason("Email must not be empty")
	}
	if len(u.PasswordHash) == 0 {
		err.AddReason("Password must not be empty")
	}

	if len(err.Reasons) > 0 {
		return err
	} else {
		return nil
	}
}

//
// Private Methods
//

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashed), err
}

func passwordIsCorrect(hash string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}

//
// Data Access Methods (no real business logic or validation
//

func getUserById(db *sqlx.DB, id int) (*User, error) {
	user := &User{}
	err := db.QueryRowx("SELECT * FROM users WHERE id = $1", id).StructScan(user)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	return user, dbError(err)
}

func getUserByName(db *sqlx.DB, name string) (*User, error) {
	user := &User{}
	err := db.QueryRowx("SELECT * FROM users WHERE name = $1", name).StructScan(user)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	return user, dbError(err)
}

func insertUser(db *sqlx.DB, name string, email string, password string) (*User, error) {
	user := &User{}
	err := db.QueryRowx("INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3) RETURNING *", name, email, password).StructScan(user)
	if dbErr := dbError(err); dbErr != nil {
		return nil, dbErr
	} else {
		return user, nil
	}
}
