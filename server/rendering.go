package server

import (
	"encoding/json"
	"net/http"
)

// Default error implementation, used for rendering in json
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(code int, message string) Error {
	return Error{Code: code, Message: message}
}

type Meta struct {
	Code   int     `json:"code"`
	Errors []Error `json:"errors,omitempty"`
}

type Result struct {
	Meta Meta        `json:"meta"`
	Body interface{} `json:"body,omitempty"`
}

type Renderer interface {
	Result(w http.ResponseWriter, code int, result interface{})
	Error(w http.ResponseWriter, code int, errors ...Error)
}

// Default renderer type
type jsonRenderer struct{}

func DefaultRenderer() Renderer {
	return jsonRenderer{}
}

func (r jsonRenderer) Result(w http.ResponseWriter, code int, result interface{}) {
	obj := Result{
		Meta: Meta{
			Code:   code,
			Errors: nil,
		},
		Body: result,
	}

	bytes, err := json.Marshal(obj)
	if err != nil {
		renderServerError(w)
	} else {
		w.WriteHeader(code)
		w.Write(bytes)
	}
}

func (r jsonRenderer) Error(w http.ResponseWriter, code int, errors ...Error) {
	result := Result{
		Meta: Meta{
			Code:   code,
			Errors: errors,
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
					"domain": "RenderError",
					"code": 500,
					"message": "Failed to generate JSON"
				}
			]
		}
	}`))
}
