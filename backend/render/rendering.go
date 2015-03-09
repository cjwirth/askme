package render

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Meta struct {
	Code   int     `json:"code"`
	Errors []Error `json:"errors,omitempty"`
}

type Result struct {
	Meta Meta        `json:"meta"`
	Body interface{} `json:"body,omitempty"`
}

// TODO: should this really be on http.ResponseWriter?
func Render(w http.ResponseWriter, code int, result interface{}) {
	r := Result{
		Meta: Meta{
			Code:   code,
			Errors: nil,
		},
		Body: result,
	}

	bytes, err := json.Marshal(r)
	if err != nil {
		RenderServerError(w)
	} else {
		w.WriteHeader(code)
		w.Write(bytes)
	}
}

func RenderError(w http.ResponseWriter, code int, errors ...Error) {
	r := Result{
		Meta: Meta{
			Code:   code,
			Errors: errors,
		},
	}

	bytes, err := json.Marshal(r)
	if err != nil {
		RenderServerError(w)
	} else {
		w.WriteHeader(code)
		w.Write(bytes)
	}
}

func RenderServerError(w http.ResponseWriter) {
	w.WriteHeader(500)
	w.Write([]byte(`{
		"meta": {
			"code": 500,
			"errors": [
				{
					"code": 500,
					"message": "Failed to generate JSON"
				}
			]
		}
	}`))
}
