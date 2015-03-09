package handling

import (
	"net/http"

	"../env"
	"../models"
	"../render"
)

func CreateUser(w http.ResponseWriter, r *http.Request, c *env.Context) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := models.InsertUser(c.DB.DB, name, email, password)

	if err != nil {
		render.RenderError(w, 500, render.Error{Code: 500, Message: err.Error()})
	} else {
		render.Render(w, 200, user)
	}
}
