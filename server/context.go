package server

import (
	"encoding/json"
)

type Context struct {
	DB      *Database
	Render  Renderer
	Decoder *json.Decoder

	PathParams map[string]string
}

func NewContext(config Config) *Context {
	c := &Context{}

	// Database
	c.DB = NewDatabase(config.DBDriver, config.DBDataSource)
	c.Render = DefaultRenderer()

	return c
}
