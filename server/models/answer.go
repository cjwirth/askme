package models

import (
	"time"

	"github.com/jmoiron/sqlx"
	sq "github.com/lann/squirrel"
)

type Answer struct {
	Id         int       `db:"id" json:"id"`
	AuthorId   int       `db:"author_id" json:"author_id"`
	QuestionId int       `db:"question_id" json:"question_id"`
	Message    string    `db:"message" json:"message"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}

func GetAnswersForQuestionId(db *sqlx.DB, qid int) ([]Answer, error) {
	as := []Answer{}
	var err error

	qry := sq.Select("*").From("answers")
	qry = qry.Where("question_id = ?", qid)
	qry = qry.OrderBy("created_at DESC")
	qry = qry.PlaceholderFormat(sq.Dollar)

	sql, params, err := qry.ToSql()
	if err != nil {
		return as, err
	}

	err = db.Select(&as, sql, params...)
	dbErr := dbError(err)
	return as, dbErr
}

func InsertAnswer(db *sqlx.DB, userId int, questionId int, message string) (*Answer, error) {
	a := Answer{}
	a.AuthorId = userId
	a.QuestionId = questionId
	a.Message = message
	var err error

	if err = ValidateAnswer(a); err != nil {
		return nil, err
	}

	qry := sq.Insert("answers").
		Columns("author_id", "question_id", "message").
		Values(a.AuthorId, a.QuestionId, a.Message).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar)

	sql, params, err := qry.ToSql()
	if err != nil {
		return nil, err
	}

	err = db.Get(&a, sql, params...)
	dbErr := dbError(err)
	if dbErr != nil {
		return nil, dbErr
	} else {
		return &a, nil
	}
}

//
// Validation
//

func ValidateAnswer(a Answer) error {
	err := NewValidationError()

	if a.AuthorId <= 0 {
		err.AddReason("Invalid author_id.")
	}
	if a.QuestionId <= 0 {
		err.AddReason("Invalid question_id.")
	}
	if a.Message == "" {
		err.AddReason("Answer must not be empty.")
	}

	if len(err.Reasons) == 0 {
		return nil
	} else {
		return err
	}
}
