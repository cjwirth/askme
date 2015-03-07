package backend

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Router(config Config) http.Handler {
	// Actual router
	router := mux.NewRouter()

	// Dependencies for handlers
	context := NewContext(config)
	common := NewChain(context)
	common.Add(LogRequest)

	// Set up routing
	router.Handle("/", common.Then(Root))

	return router
}

func Root(w http.ResponseWriter, r *http.Request, c *Context) {

	fmt.Println("Root Handler")
	w.WriteHeader(200)
	fmt.Fprint(w, "Root")
}
