package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nonameb3/golang-mux-api/model"
)

// GetBooks function
func GetBooks(w http.ResponseWriter, r *http.Request) {
	result := ResponseList{false, "", model.Books}
	result.responseBody(r.Header.Get("Content-type"), w)
}

// GetBook function
func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Gets params
	var book *model.Book
	for _, b := range model.Books {
		if b.ID == params["ID"] {
			book = &b
			break
		}
	}

	result := Response{false, "", book}
	result.responseBody(r.Header.Get("Content-type"), w)
}

// CreateBook function
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook model.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		result := Response{true, err.Error(), nil}
		result.responseBody(r.Header.Get("Content-type"), w)
		return
	}

	newBook.ID = strconv.Itoa(rand.Intn(100000)) // Mock ID - not safe
	model.Books = append(model.Books, newBook)

	result := Response{false, "", &newBook}
	result.responseBody(r.Header.Get("Content-type"), w)
}

// UpdateBook function
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Gets params
	for idx, b := range model.Books {
		if b.ID == params["ID"] {
			var newBook model.Book
			_ = json.NewDecoder(r.Body).Decode(&newBook)
			newBook.ID = params["ID"]
			model.Books[idx] = newBook
			result := Response{false, "", &newBook}
			result.responseBody(r.Header.Get("Content-type"), w)
			return
		}
	}
	result := Response{false, "", nil}
	result.responseBody(r.Header.Get("Content-type"), w)
}

// DeleteBook function
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Gets params
	for idx, b := range model.Books {
		if b.ID == params["ID"] {
			model.Books = append(model.Books[:idx], model.Books[idx+1:]...)
		}
	}
	result := ResponseList{false, "", model.Books}
	result.responseBody(r.Header.Get("Content-type"), w)
}

// SetString function
func SetString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

// GetJSON function
func GetJSON(w http.ResponseWriter, r *http.Request) {
	result := Response{false, "Hello there!", nil}
	result.responseBody(r.Header.Get("Content-type"), w)
}
