package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"askme/server"
	"github.com/gorilla/mux"
)

// InitContext is the first middleware to be called. It will initialize the context with
// request-specific data and objects that might be helpful in handling the request
func InitContext(w http.ResponseWriter, r *http.Request, c *server.Context, next server.Handler) {
	c.Decoder = json.NewDecoder(r.Body)
	c.PathParams = mux.Vars(r)
	next(w, r, c)
}

func LogRequest(w http.ResponseWriter, r *http.Request, c *server.Context, next server.Handler) {
	log.Println("Request: " + r.Method + " " + r.RequestURI)
	next(w, r, c)
}
