package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"askme/server"
	"askme/server/models"
)

type AnswerParam struct {
	QuestionId float64 `json:"question_id"`
	Message    string  `json:"message"`
}

func QueryAnswers(w http.ResponseWriter, r *http.Request, c *server.Context) {
	questionId, _ := strconv.Atoi(r.FormValue("question_id"))

	as, err := models.GetAnswersForQuestionId(c.DB.DB, questionId)
	if err != nil {
		c.Render.ServerError(w, err)
	} else {
		c.Render.ResultOK(w, as)
	}
}

func CreateAnswer(w http.ResponseWriter, r *http.Request, c *server.Context) {
	userId := c.Session.UserId
	var a AnswerParam
	if c.MustDecodeBody(w, &a) == false {
		return
	}
	questionId := int(a.QuestionId)

	if models.GetQuestionById(c.DB.DB, questionId) == nil {
		c.Render.BadRequest(w, errors.New("Question does not exist"))
		return
	}

	answer, err := models.InsertAnswer(c.DB.DB, userId, questionId, a.Message)
	if err != nil {
		c.Render.BadRequest(w, err)
	} else {
		c.Render.ResultOK(w, answer)
	}
}
