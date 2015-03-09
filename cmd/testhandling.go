package main

import (
	"net/http"

	_ "github.com/lib/pq"

	"../backend/env"
	"../backend/handling"
)

// Here begin tests

func main() {
	config := env.Config{
		DBDriver:     "postgres",
		DBDataSource: "user=askme dbname=askme_dev",
	}
	r := handling.Router(config)
	http.ListenAndServe(":8000", r)
}
