package model

// Init books var as a slice Book struct
var books []Book

// Book struct
type Book struct {
	ID     string  `json:"id"`
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
