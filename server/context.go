package server

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Context struct {
	DB      *Database
	Render  Renderer
	Decoder *json.Decoder
	Session *Session

	PathParams map[string]string
}

func NewContext(config Config) *Context {
	c := &Context{}

	// Database
	c.DB = NewDatabase(config.DBDriver, config.DBDataSource)
	c.Render = DefaultRenderer(config)

	return c
}

func (c *Context) MustDecodeBody(w http.ResponseWriter, v interface{}) bool {
	if err := c.Decoder.Decode(&v); err != nil {
		c.Render.BadRequest(w, errors.New("Could not decode input"))
		return false
	}
	return true
}
