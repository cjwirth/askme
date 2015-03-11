package server

import (
	"encoding/json"
	"net/http"
)

type ServerError struct {
	Message string `json:"message"`
}

func makeErrors(errors ...error) []ServerError {
	serverErrors := []ServerError{}
	for _, err := range errors {
		serverErrors = append(serverErrors, ServerError{Message: err.Error()})
	}
	return serverErrors
}

type meta struct {
	Code   int           `json:"code"`
	Errors []ServerError `json:"errors,omitempty"`
}

type result struct {
	Meta meta        `json:"meta"`
	Body interface{} `json:"body,omitempty"`
}

type Renderer interface {
	Result(w http.ResponseWriter, code int, result interface{})
	Error(w http.ResponseWriter, code int, errors ...error)
}

// Default renderer type
type jsonRenderer struct{}

func DefaultRenderer() Renderer {
	return jsonRenderer{}
}

func (r jsonRenderer) Result(w http.ResponseWriter, code int, payload interface{}) {
	obj := result{
		Meta: meta{
			Code:   code,
			Errors: nil,
		},
		Body: payload,
	}

	bytes, err := json.Marshal(obj)
	if err != nil {
		renderServerError(w)
	} else {
		w.WriteHeader(code)
		w.Write(bytes)
	}
}

func (r jsonRenderer) Error(w http.ResponseWriter, code int, errors ...error) {
	result := result{
		Meta: meta{
			Code:   code,
			Errors: makeErrors(errors...),
		},
	}

	bytes, err := json.Marshal(result)
	if err != nil {
		renderServerError(w)
	} else {
		w.WriteHeader(code)
		w.Write(bytes)
	}
}

func renderServerError(w http.ResponseWriter) {
	w.WriteHeader(500)
	w.Write([]byte(`{
		"meta": {
			"code": 500,
			"errors": [
				{
					"message": "Failed to generate JSON"
				}
			]
		}
	}`))
}
