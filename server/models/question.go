package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	sq "github.com/lann/squirrel"
)

type Question struct {
	Id        int       `db:"id" json:"id"`
	AuthorId  int       `db:"author_id" json:"author_id"`
	Title     string    `db:"title" json:"title"`
	Question  string    `db:"question" json:"question"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func QueryQuestions(db *sqlx.DB, authorId int, query string, offset int) ([]Question, error) {
	qs := []Question{}
	qry := sq.Select("*").From("questions")

	if authorId != 0 {
		qry = qry.Where("author_id = ?", authorId)
	}
	if query != "" {
		word := fmt.Sprint("%", query, "%")
		qry = qry.Where("(title LIKE ? OR question LIKE ?)", word, word)
	}
	if offset > 0 {
		qry = qry.Offset(uint64(offset))
	}

	qry = qry.OrderBy("created_at DESC")
	qry = qry.PlaceholderFormat(sq.Dollar)
	sql, params, err := qry.ToSql()

	if err != nil {
		return qs, err
	} else {
		err := db.Select(&qs, sql, params...)
		dbErr := dbError(err)
		return qs, dbErr
	}
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
