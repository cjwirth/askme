package handling

import (
	"encoding/json"
	"log"
	"net/http"

	"../env"
)

// InitContext is the first middleware to be called. It will initialize the context with
// request-specific data and objects that might be helpful in handling the request
func InitContext(w http.ResponseWriter, r *http.Request, c *env.Context, next Handler) {
	c.Decoder = json.NewDecoder(r.Body)
	next(w, r, c)
}

func LogRequest(w http.ResponseWriter, r *http.Request, c *env.Context, next Handler) {
	log.Println("Request: " + r.Method + " " + r.RequestURI)
	next(w, r, c)
}
