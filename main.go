package main

import (
	"encoding/json"
	"log"
	"net/http"

	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
)

// Bok struct (Model)
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`

}

//
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get All books
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single Book.
func getBook(w http.ResponseWriter, r *http.Request){
	
}

// Create new Book
func createBook(w http.ResponseWriter, r *http.Request){
	
}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request){
	
}
func deleteBook(w http.ResponseWriter, r *http.Request){
	
}



func main(){
	// Init Router
	r := mux.NewRouter()
	
	// mock Data
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book one", Author: &Author{
		Firstname: "John", Lastname: "Doe",
	}})
	books = append(books, Book{ID: "2", Isbn: "448734", Title: "Book two", Author: &Author{
		Firstname: "John", Lastname: "Doe",
	}})


	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))


}