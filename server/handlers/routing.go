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
	common.Add(InitContext, RestoreSession, LogRequest)

	// Router for handlers requiring login
	login := common.Branch(RequireLogin)

	// Set up routing
	router.Handle("/", common.Then(Root))

	// Users
	router.Handle("/users", common.Then(CreateUser)).Methods("POST")
	router.Handle("/users/{id:[0-9]+}", common.Then(GetUser)).Methods("GET")
	router.Handle("/users/me", login.Then(GetUserMe)).Methods("GET")

	// Login
	router.Handle("/login", common.Then(Login)).Methods("POST")
	router.Handle("/logout", common.Then(Logout)).Methods("POST")

	return router
}

func Root(w http.ResponseWriter, r *http.Request, c *server.Context) {

	fmt.Println("Root Handler")
	w.WriteHeader(200)
	fmt.Fprint(w, "Root")
}
