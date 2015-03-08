package handling

import (
	"log"
	"net/http"

	"../env"
)

func LogRequest(w http.ResponseWriter, r *http.Request, c *env.Context, next Handler) {
	log.Println("Request: " + r.Method + " " + r.RequestURI)
	next(w, r, c)
}
