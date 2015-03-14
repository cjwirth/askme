package server

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ServerError struct {
	Message string `json:"message"`
}

type meta struct {
	Status int          `json:"status"`
	Error  *ServerError `json:"error,omitempty"`
}

type result struct {
	Meta meta        `json:"meta"`
	Body interface{} `json:"body,omitempty"`
}

type Renderer interface {

	// ResultOK sends a successfull call with some payback to the http response
	ResultOK(w http.ResponseWriter, result interface{})

	// ResultStatus sends a successful request, with a custom http status
	// For instance, if a post is successfully created, can return
	// http.StatusCreated instead of http.StatusOK
	ResultStatus(w http.ResponseWriter, status int, result interface{})

	// NotFound sends a not found response
	NotFound(w http.ResponseWriter)

	// BadRequest sends a bad request response
	// TODO: figure out what the error should actually do
	BadRequest(w http.ResponseWriter, err error)

	// UnknownServerError writes a standard unknown error to the http response
	UnknownServerError(w http.ResponseWriter)

	// ServerError writes a server error to the http response
	// It also passes in the error that occurred
	ServerError(w http.ResponseWriter, err error)

	// Error sends an error to the http response
	// status is the HTTP Status to send
	// err is the error to send
	Error(w http.ResponseWriter, status int, err error)

	// Write writes to the http response with the given status and payload
	// status is the HTTP status
	// err is what error is passed in as the error
	// payload is what is filled in as the body
	Write(w http.ResponseWriter, status int, payload interface{}, err error)
}

// Default renderer type
type jsonRenderer struct {
	prettyPrint bool
}

func DefaultRenderer(c Config) Renderer {
	return jsonRenderer{prettyPrint: c.JSONPrettyPrint}
}

func (r jsonRenderer) ResultOK(w http.ResponseWriter, payload interface{}) {
	r.Write(w, http.StatusOK, payload, nil)
}

func (r jsonRenderer) ResultStatus(w http.ResponseWriter, status int, result interface{}) {
	r.Write(w, status, result, nil)
}

func (r jsonRenderer) NotFound(w http.ResponseWriter) {
	r.Write(w, http.StatusNotFound, nil, errors.New("Not Found"))
}

func (r jsonRenderer) BadRequest(w http.ResponseWriter, err error) {
	r.Write(w, http.StatusBadRequest, nil, err)
}

func (r jsonRenderer) UnknownServerError(w http.ResponseWriter) {
	r.ServerError(w, errors.New("Request could not be completed"))
}

func (r jsonRenderer) ServerError(w http.ResponseWriter, err error) {
	r.Error(w, http.StatusInternalServerError, err)
}

func (r jsonRenderer) Error(w http.ResponseWriter, status int, err error) {
	r.Write(w, status, nil, err)
}

func (r jsonRenderer) Write(w http.ResponseWriter, status int, payload interface{}, err error) {
	obj := result{Meta: meta{}}
	obj.Meta.Status = status
	if err != nil {
		obj.Meta.Error = &ServerError{Message: err.Error()}
	}
	if payload != nil {
		obj.Body = payload
	}

	var bytes []byte
	var encErr error
	if r.prettyPrint {
		bytes, encErr = json.MarshalIndent(obj, "", "\t")
	} else {
		bytes, encErr = json.Marshal(obj)
	}
	if encErr != nil {
		renderServerError(w)
	} else {
		w.WriteHeader(status)
		w.Write(bytes)
	}
}

// renderServerError writes a standard 500 error
// Used if the JSON data cannot be generated
func renderServerError(w http.ResponseWriter) {
	w.WriteHeader(500)
	w.Write([]byte(`{
		"meta": {
			"status": 500,
			"error":
				{
					"message": "Failed to generate JSON"
				}
		}
	}`))
}
