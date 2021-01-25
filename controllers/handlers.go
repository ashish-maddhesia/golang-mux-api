package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getBooks(w http.ResponseWriter, r *http.Request) {
	result := ResponseList{false, "", books}
	result.responseBody(r.Header.Get("Content-type"), w)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Gets params
	var book *Book
	for _, b := range books {
		if b.ID == params["ID"] {
			book = &b
			break
		}
	}

	result := Response{false, "", book}
	result.responseBody(r.Header.Get("Content-type"), w)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		result := Response{true, err.Error(), nil}
		result.responseBody(r.Header.Get("Content-type"), w)
		return
	}

	newBook.ID = strconv.Itoa(rand.Intn(100000)) // Mock ID - not safe
	books = append(books, newBook)

	result := Response{false, "", &newBook}
	result.responseBody(r.Header.Get("Content-type"), w)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Gets params
	for idx, b := range books {
		if b.ID == params["ID"] {
			var newBook Book
			_ = json.NewDecoder(r.Body).Decode(&newBook)
			newBook.ID = params["ID"]
			books[idx] = newBook
			result := Response{false, "", &newBook}
			result.responseBody(r.Header.Get("Content-type"), w)
			return
		}
	}
	result := Response{false, "", nil}
	result.responseBody(r.Header.Get("Content-type"), w)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Gets params
	for idx, b := range books {
		if b.ID == params["ID"] {
			books = append(books[:idx], books[idx+1:]...)
		}
	}
	result := ResponseList{false, "", books}
	result.responseBody(r.Header.Get("Content-type"), w)
}

func setString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func getJSON(w http.ResponseWriter, r *http.Request) {
	result := Response{false, "Hello there!", nil}
	result.responseBody(r.Header.Get("Content-type"), w)
}
