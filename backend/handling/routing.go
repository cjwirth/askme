package handling

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"../env"
)

func Router(config env.Config) http.Handler {
	// Actual router
	router := mux.NewRouter()

	// Dependencies for handlers
	context := env.NewContext(config)
	common := NewChain(context)
	common.Add(LogRequest)

	// Set up routing
	router.Handle("/", common.Then(Root))

	return router
}

func Root(w http.ResponseWriter, r *http.Request, c *env.Context) {

	fmt.Println("Root Handler")
	w.WriteHeader(200)
	fmt.Fprint(w, "Root")
}
