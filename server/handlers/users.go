package handlers

import (
	"errors"
	"net/http"

	"askme/server"
	"askme/server/models"
)

// UserParam is the form that user objects should come in as JSON
type UserParam struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUser(w http.ResponseWriter, r *http.Request, c *server.Context) {
	userId := c.PathParams["id"]

	user := models.GetUserById(c.DB.DB, userId)
	if user != nil {
		c.Render.ResultOK(w, user)
	} else {
		c.Render.NotFound(w)
	}
}

// CreateUser is an http endpoint handler
// Path: POST /user
// Param: HTTP Body is a UserParam in JSON format
func CreateUser(w http.ResponseWriter, r *http.Request, c *server.Context) {
	var u UserParam
	if err := c.Decoder.Decode(&u); err != nil {
		c.Render.BadRequest(w, errors.New("Could not decode inputn"))
		return
	}

	user, err := models.InsertUser(c.DB.DB, u.Name, u.Email, u.Password)

	c.Render.Result(w, user, err)
}
