package main

	

import (
	"encoding/json"
	"log"
	"net/http"
	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
)

// Book Struct(Model)
type Book struct {
	ID 		string	`json:"id"`
	Isbn	string	`json:"isbn"`
	Title	string	`json:"title"`
	Author	*Author `json:"author"`
}

//Author struct
type Author struct {
	Firstname	string	`json:"firstname"`
	Lastname	string	`json:"lastname"`
}

// Init books var as a slice book struct
var books []Book

//get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
//get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	//
}
//create new book
func createBook(w http.ResponseWriter, r *http.Request) {
	//
}
//get All Books
func updateBook(w http.ResponseWriter, r *http.Request) {
	//
}
//get All Books
func deleteBook(w http.ResponseWriter, r *http.Request) {
	//
}

func main() {
	//init router
	r:= mux.NewRouter()

	//mock Data - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "436455", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "436444", Title: "Book two", Author: &Author{Firstname: "Samuel", Lastname: "Etoo"}})

	//Route Handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
