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

// CreateUser is an http endpoint handler
// Path: POST /user
// Param: HTTP Body is a UserParam in JSON format
func CreateUser(w http.ResponseWriter, r *http.Request, c *server.Context) {
	var u UserParam
	if err := c.Decoder.Decode(&u); err != nil {
		c.Render.Error(w, http.StatusBadRequest, errors.New("Could not decode input"))
		return
	}

	user, errs := models.InsertUser(c.DB.DB, u.Name, u.Email, u.Password)

	if len(errs) > 0 {
		c.Render.Error(w, http.StatusBadRequest, errs...)
	} else {
		c.Render.Result(w, http.StatusOK, user)
	}
}
