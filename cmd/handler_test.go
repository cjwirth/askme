package main

import (
	"fmt"
	"net/http"

	"../backend"

	"github.com/gorilla/mux"
)

// Here begin tests

func LogContext(w http.ResponseWriter, r *http.Request, c *backend.Context, next backend.Handler) {
	fmt.Println("Context: " + c.Name)
	next(w, r, c)
}

func LogRequest(w http.ResponseWriter, r *http.Request, c *backend.Context, next backend.Handler) {
	fmt.Println("Request: " + r.Method + r.RequestURI)
	next(w, r, c)
}

func LogAdmin(w http.ResponseWriter, r *http.Request, c *backend.Context, next backend.Handler) {
	fmt.Println("Admin Checker")
	next(w, r, c)
}

func Root(w http.ResponseWriter, r *http.Request, c *backend.Context) {
	fmt.Println("Root Handler")
	w.WriteHeader(200)
	fmt.Fprint(w, "Root")
}

func main() {
	context := &backend.Context{"Steve"}
	chain := backend.NewChain(context).Add(LogContext, LogRequest)
	admin := chain.Branch(LogAdmin)

	r := mux.NewRouter()
	r.Handle("/", chain.Then(Root))
	r.Handle("/admin", admin.Then(Root))
	http.ListenAndServe(":4000", r)
}
