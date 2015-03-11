package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"askme/server"
	"askme/server/handlers"
)

// Here begin tests

func main() {
	config := server.Config{
		DBDriver:     "postgres",
		DBDataSource: "user=askme dbname=askme_dev",
	}
	r := handlers.Router(config)
	fmt.Println(http.ListenAndServe(":8000", r))
}
