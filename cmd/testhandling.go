package main

import (
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"../backend/env"
	"../backend/handling"
)

// Here begin tests

func main() {
	config := env.Config{
		DBDriver:     "sqlite3",
		DBDataSource: "file:develop.db",
	}
	r := handling.Router(config)
	http.ListenAndServe(":8000", r)
}
