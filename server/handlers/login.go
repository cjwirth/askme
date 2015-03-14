package handlers

import (
	"errors"
	"net/http"

	"askme/server"
	"askme/server/models"
)

type LoginParams struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request, c *server.Context) {
	var p LoginParams
	if err := c.Decoder.Decode(&p); err != nil {
		c.Render.BadRequest(w, errors.New("Could not decode input"))
		return
	}

	user := *models.GetUserByName(c.DB.DB, p.Name)
	if user.CheckPassword(p.Password) {
		session := server.NewSession()
		session.UserId = user.Id
		server.SetSession(w, session)
		c.Render.ResultOK(w, nil)
	} else {
		c.Render.BadRequest(w, nil)
	}
}

func Logout(w http.ResponseWriter, r *http.Request, c *server.Context) {
	server.DeleteSession(w, c.Session)
	c.Render.ResultOK(w, nil)
}
