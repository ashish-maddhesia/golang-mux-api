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
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	var book *model.Book
	for _, b := range model.Books {
		if b.ID == params["ID"] {
			book = &b
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	http.Error(w," book not found",http.StatusNotFound)
	
}

// CreateBook function
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBook model.Book
	_ := json.NewDecoder(r.Body).Decode(&newBook)
	newBook.ID = strconv.Itoa(rand.Intn(100000)) // Mock ID - not safe
	model.Books = append(model.Books, newBook)
    json.NewEncoder(w).Encode(newBook)
	Fprintln(newBook. time.Now())
}

// UpdateBook function
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	for idx, b := range model.Books {
		if b.ID == params["ID"] {
			var newBook model.Book
			_ = json.NewDecoder(r.Body).Decode(&newBook)
			newBook.ID = params["ID"]
			model.Books[idx] = newBook
			result := Response{false, "", &newBook}
			Fprintln(result.time.Now())
			return
		}
	}
	
	
}

// DeleteBook function
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
    for idx, b := range model.Books {
		if b.ID == params["ID"] {
			model.Books = append(model.Books[:idx], model.Books[idx+1:]...)
			return
		}
	}
	json.NewEncoder(w).Encode(b)
	
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
