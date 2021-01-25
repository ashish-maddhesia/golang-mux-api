package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/nonameb3/golang-mux-api/model"
)

// Response struct
type Response struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    *model.Book `json:"data"`
}

// ResponseList struct
type ResponseList struct {
	Error   bool         `json:"error"`
	Message string       `json:"message"`
	Data    []model.Book `json:"data"`
}

func (r Response) responseBody(contentType string, w http.ResponseWriter) {
	const JSON = "application/json"
	switch contentType {
	default:
		w.Header().Set("Content-Type", JSON)
		json.NewEncoder(w).Encode(r) //asked for json, return json
	}
}

func (r ResponseList) responseBody(contentType string, w http.ResponseWriter) {
	const JSON = "application/json"
	switch contentType {
	default:
		w.Header().Set("Content-Type", JSON)
		json.NewEncoder(w).Encode(r) //asked for json, return json
	}
}
