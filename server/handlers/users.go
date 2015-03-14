package handlers

import (
	"net/http"
	"strconv"

	"askme/server"
	"askme/server/models"
)

// UserParam is the form that user objects should come in as JSON
type UserParam struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUserMe(w http.ResponseWriter, r *http.Request, c *server.Context) {
	id := c.Session.UserId
	user := models.GetUserById(c.DB.DB, id)
	if user != nil {
		c.Render.ResultOK(w, user)
	} else {
		c.Render.NotFound(w)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request, c *server.Context) {
	id := c.PathParams["id"]
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.Render.NotFound(w)
		return
	}

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
	if c.MustDecodeBody(w, &u) == false {
		return
	}

	user, err := models.InsertUser(c.DB.DB, u.Name, u.Email, u.Password)
	if err != nil {
		c.Render.BadRequest(w, err)
	} else {
		c.Render.ResultOK(w, user)
	}
}
