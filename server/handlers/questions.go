package handlers

import (
	"net/http"
	"strconv"

	"askme/server"
	"askme/server/models"
)

type QuestionParam struct {
	Title    string `json:"title"`
	Question string `json:"question"`
}

func GetQuestion(w http.ResponseWriter, r *http.Request, c *server.Context) {
	id := c.PathParams["id"]
	questionId, err := strconv.Atoi(id)
	if err != nil {
		c.Render.NotFound(w)
		return
	}

	question := models.GetQuestionById(c.DB.DB, questionId)
	if question != nil {
		c.Render.ResultOK(w, question)
	} else {
		c.Render.NotFound(w)
	}
}

func CreateQuestion(w http.ResponseWriter, r *http.Request, c *server.Context) {
	userId := c.Session.UserId
	var q QuestionParam
	if c.MustDecodeBody(w, &q) == false {
		return
	}

	question, err := models.InsertQuestion(c.DB.DB, userId, q.Title, q.Question)
	if err != nil {
		c.Render.BadRequest(w, err)
	} else {
		c.Render.ResultOK(w, question)
	}
}
