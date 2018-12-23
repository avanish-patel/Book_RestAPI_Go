package main

import (
	"strconv"
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"github.com/gorilla/mux"
)


// Book Struct (Model)
type Book struct {
	ID	string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`

}

// Author Struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}

// Get book by Id
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	// Get url param from request
	param := mux.Vars(r)
	
	// loop through the books and find by id
	for _, book := range books {
		if book.ID == param["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Create Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000)) // Not safe 
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}


// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Init the router
	r := mux.NewRouter()

	// Mock Data
	books = append(books, Book{ID:"1",Isbn:"448743", Title:"Book One", Author: &Author{FirstName:"Avanish", LastName:"Patel"}})
	books = append(books, Book{ID:"2",Isbn:"448744", Title:"Book Two", Author: &Author{FirstName:"Dixit", LastName:"Patel"}})
	books = append(books, Book{ID:"3",Isbn:"448745", Title:"Book Three", Author: &Author{FirstName:"Tejas", LastName:"Patel"}})



	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("api/books/{id}", deleteBook).Methods("DELETE")

	// Run server
	log.Fatal(http.ListenAndServe(":8080", r))

}