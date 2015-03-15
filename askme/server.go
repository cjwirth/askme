package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"askme/server"
	"askme/server/handlers"
)

func main() {
	config := server.Config{
		DBDriver:        "postgres",
		DBDataSource:    "user=askme dbname=askme_dev",
		JSONPrettyPrint: true,
	}
	r := handlers.Router(config)
	fmt.Println(http.ListenAndServe(":8000", r))
}
