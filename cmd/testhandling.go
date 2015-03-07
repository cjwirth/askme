package main

import (
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"../backend"
)

// Here begin tests

func main() {
	config := backend.Config{
		DBDriver:     "sqlite3",
		DBDataSource: "file:develop.db",
	}
	r := backend.Router(config)
	http.ListenAndServe(":8000", r)
}
