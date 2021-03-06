package handlers

import (
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

	// Users
	router.Handle("/users", common.Then(CreateUser)).Methods("POST")
	router.Handle("/users/{id:[0-9]+}", common.Then(GetUser)).Methods("GET")
	router.Handle("/users/me", login.Then(GetUserMe)).Methods("GET")

	// Questions
	router.Handle("/questions", common.Then(QueryQuestions)).Methods("GET")
	router.Handle("/questions/{id:[0-9]+}", common.Then(GetQuestion)).Methods("GET")
	router.Handle("/questions", login.Then(CreateQuestion)).Methods("POST")

	router.Handle("/answers", common.Then(QueryAnswers)).Methods("GET")
	router.Handle("/answers", login.Then(CreateAnswer)).Methods("POST")

	// Login
	router.Handle("/login", common.Then(Login)).Methods("POST")
	router.Handle("/logout", common.Then(Logout)).Methods("POST")

	router.NotFoundHandler = common.Then(NotFound)

	return router
}

func NotFound(w http.ResponseWriter, r *http.Request, c *server.Context) {
	c.Render.NotFound(w)
}
