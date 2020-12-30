package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// inital route
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	books = append(books, Book{ID: "1", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route handles & endpoints
	r.HandleFunc("/", setString).Methods("GET")
	r.HandleFunc("/api/json", getJSON).Methods("GET")
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{ID}", getBook).Methods("GET")
	r.HandleFunc("/api/books/{ID}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{ID}", deleteBook).Methods("DELETE")

	// Start server
	port := ":10000"
	log.Println("listen on", port)
	log.Fatal(http.ListenAndServe(port, r))
}
