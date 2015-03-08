package env

import ()

type Context struct {
	DB *Database
}

func NewContext(config Config) *Context {
	c := &Context{}

	// Database
	c.DB = NewDatabase(config.DBDriver, config.DBDataSource)

	return c
}
