package main

import (
	"encoding/json"
	"net/http"
)

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
