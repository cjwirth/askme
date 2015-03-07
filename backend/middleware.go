package backend

import (
	"log"
	"net/http"
)

func LogRequest(w http.ResponseWriter, r *http.Request, c *Context, next Handler) {
	log.Println("Request: " + r.Method + " " + r.RequestURI)
	next(w, r, c)
}
