package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"askme/server"
)

func Router(config server.Config) http.Handler {
	// Actual router
	router := mux.NewRouter()

	// Dependencies for handlers
	context := server.NewContext(config)
	common := server.NewChain(context)
	common.Add(InitContext, LogRequest)

	// Set up routing
	router.Handle("/", common.Then(Root))
	router.Handle("/users", common.Then(CreateUser)).Methods("POST")

	return router
}

func Root(w http.ResponseWriter, r *http.Request, c *server.Context) {

	fmt.Println("Root Handler")
	w.WriteHeader(200)
	fmt.Fprint(w, "Root")
}
