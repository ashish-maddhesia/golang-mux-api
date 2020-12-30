package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Init books var as a slice Book struct
var books []Book

// Book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Response struct
type Response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    *Book  `json:"data"`
}

// ResponseList struct
type ResponseList struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    []Book `json:"data"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	result := ResponseList{false, "", books}
	responseBodyList(result, r.Header.Get("Content-type"), w)
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
	responseBody(result, r.Header.Get("Content-type"), w)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w)
}

func setString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func getJSON(w http.ResponseWriter, r *http.Request) {
	result := Response{false, "Hello there!", nil}
	responseBody(result, r.Header.Get("Content-type"), w)
}

func responseBody(r Response, contentType string, w http.ResponseWriter) {
	const JSON = "application/json"
	switch contentType {
	default:
		w.Header().Set("Content-Type", JSON)
		json.NewEncoder(w).Encode(r) //asked for json, return json
	}
}

func responseBodyList(r ResponseList, contentType string, w http.ResponseWriter) {
	const JSON = "application/json"
	switch contentType {
	default:
		w.Header().Set("Content-Type", JSON)
		json.NewEncoder(w).Encode(r) //asked for json, return json
	}
}

func main() {
	// inital route
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route handles & endpoints
	r.HandleFunc("/", setString).Methods("GET")
	r.HandleFunc("/api/json", getJSON).Methods("GET")
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{ID}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{ID}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{ID}", deleteBook).Methods("DELETE")

	// Start server
	port := ":10000"
	log.Println("listen on", port)
	log.Fatal(http.ListenAndServe(port, r))
}
