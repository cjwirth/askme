package models

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type Question struct {
	Id        int       `db:"id" json:"id"`
	AuthorId  int       `db:"author_id" json:"author_id"`
	Title     string    `db:"title" json:"title"`
	Question  string    `db:"question" json:"question"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func GetQuestionById(db *sqlx.DB, id int) *Question {
	q, _ := getQuestionById(db, id)
	return q
}

func InsertQuestion(db *sqlx.DB, authorId int, title string, question string) (*Question, error) {
	q := Question{}
	q.AuthorId = authorId
	q.Title = title
	q.Question = question

	if err := ValidateQuestion(q); err != nil {
		return nil, err
	}

	return insertQuestion(db, authorId, title, question)
}

//
// Validation
//

func ValidateQuestion(q Question) error {
	err := NewValidationError()

	if q.AuthorId <= 0 {
		err.AddReason("Invalid author_id.")
	}
	if q.Title == "" {
		err.AddReason("Title must not be empty.")
	}
	if q.Question == "" {
		err.AddReason("Question must not be empty.")
	}

	if len(err.Reasons) > 0 {
		return err
	} else {
		return nil
	}
}

//------------------------------------------------------------------------------
// Private Methods

//
// Data Access Methods
//

func getQuestionById(db *sqlx.DB, id int) (*Question, error) {
	var q Question
	err := db.QueryRowx("SELECT * FROM questions WHERE id = $1", id).StructScan(&q)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	return &q, dbError(err)
}

func insertQuestion(db *sqlx.DB, authorId int, title string, question string) (*Question, error) {
	var q Question
	err := db.QueryRowx("INSERT INTO questions (author_id, title, question) VALUES ($1, $2, $3) RETURNING *", authorId, title, question).StructScan(&q)
	if dbErr := dbError(err); dbErr != nil {
		return nil, dbErr
	} else {
		return &q, nil
	}
}
