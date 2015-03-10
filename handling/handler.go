package handling

import (
	"../env"
	"net/http"
)

// Handler is just a function that responds to http requests with the context also passed in
type Handler func(http.ResponseWriter, *http.Request, *env.Context)

// Middleware is a function that acts upon an http request, handles it, and calls the next handler
// If a middleware can also stop the chain of events by not calling the next handler
type Middleware func(http.ResponseWriter, *http.Request, *env.Context, Handler)

// node is a step in the middleware list.
type node struct {
	middleware Middleware
	next       *node
}

// emptyNode represents the end of the middleware list, and does nothing (except stops the request...)
var emptyNode = &node{func(w http.ResponseWriter, r *http.Request, c *env.Context, next Handler) {}, &node{}}

// add appends the given middlewares to the list
// It will call itself recursively to get it to the end
func (n *node) add(middlewares ...Middleware) {
	if n.next != emptyNode {
		n.next.add(middlewares...)
		return
	}

	count := len(middlewares)

	if count > 0 {
		n.next = &node{middlewares[0], emptyNode}
	}
	if count > 1 {
		n.next.add(middlewares[1:]...)
	}
}

// copy returns a copy of the node, including all the nodes after this one
// In effect, it creates a copy of the entire list
func (n node) copy() *node {
	if n.next == emptyNode {
		return &node{n.middleware, emptyNode}
	}

	next := n.next.copy()
	return &node{n.middleware, next}
}

// ServeHTTP is called when a request comes in. The node calls its middleware, and passes in the next step
func (n node) ServeHTTP(w http.ResponseWriter, r *http.Request, c *env.Context) {
	n.middleware(w, r, c, n.next.ServeHTTP)
}

// MiddlewareChain is an http.Handler that contains a list of middlewares
// Each middleware function is called in succession until at last the handler is called (or a middleware cancels it)
type MiddlewareChain struct {
	Context *env.Context
	head    *node
}

// NewChain returns a new MiddlewareChain with no actual middleware attached
func NewChain(c *env.Context) *MiddlewareChain {
	// Dummy middleware. Only exists to kickstart the thing and deal with no null pointer errors
	startingMiddleware := func(w http.ResponseWriter, r *http.Request, c *env.Context, next Handler) {
		next(w, r, c)
	}
	return &MiddlewareChain{c, &node{startingMiddleware, emptyNode}}
}

// ServeHTTP is the http.Handler implementation. Sends the request down the middleware list
func (c MiddlewareChain) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.head.ServeHTTP(w, r, c.Context)
}

// Add appends the given middleware to be called at the end of the list.
// Returns the chain itself
func (c *MiddlewareChain) Add(middlewares ...Middleware) *MiddlewareChain {
	c.head.add(middlewares...)
	return c
}

// Branch returns a copy of the chain it was called upon, and appends the given middleware to the end of the chain
func (c MiddlewareChain) Branch(middlewares ...Middleware) *MiddlewareChain {
	c.head = c.head.copy()
	c.head.add(middlewares...)
	return &c
}

// Then returns a copy of the chain it was called upon, with the handler appended to the end
// This method is meant to be used to be the final call on chains before sending them to the router, or wherever.
func (c MiddlewareChain) Then(h Handler) *MiddlewareChain {
	chain := c.Branch(func(w http.ResponseWriter, r *http.Request, c *env.Context, next Handler) {
		h(w, r, c)
		next(w, r, c)
	})
	return chain
}
