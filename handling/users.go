package handling

import (
	"net/http"

	"../env"
	"../models"
)

type UserParam struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request, c *env.Context) {
	var u UserParam
	if err := c.Decoder.Decode(&u); err != nil {
		c.Render.Error(w, http.StatusBadRequest, env.NewError(403, "Could not decode input"))
		return
	}

	user, errs := models.InsertUser(c.DB.DB, u.Name, u.Email, u.Password)
	errors := []env.Error{}
	for _, err := range errs {
		errors = append(errors, env.NewError(0, err.Error()))
	}

	if len(errors) > 0 {
		c.Render.Error(w, http.StatusBadRequest, errors...)
	} else {
		c.Render.Result(w, http.StatusOK, user)
	}
}
