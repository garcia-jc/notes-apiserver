package api

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Errors []ErrorDetail `json:"errors,omitempty"`
}

type ErrorDetail struct {
	Status int    `json:"status,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
}

func newError(status int, title, detail string) APIError {
	return APIError{
		Errors: []ErrorDetail{{
			Status: status,
			Title:  title,
			Detail: detail,
		}},
	}
}

func (a APIError) Write(w http.ResponseWriter) {
	w.WriteHeader(a.Errors[0].Status)
	_ = json.NewEncoder(w).Encode(a)
}
