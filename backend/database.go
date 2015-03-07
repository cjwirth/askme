package backend

import (
	"github.com/jmoiron/sqlx"
)

// Database is a wrapper for my data storage tools.
// Unclear as to whether it is necessary, but I will keep it around
type Database struct {
	DB *sqlx.DB
}

func NewDatabase(driver string, dataSource string) *Database {
	db := &Database{}

	db.DB = sqlx.MustOpen(driver, dataSource)

	return db
}
