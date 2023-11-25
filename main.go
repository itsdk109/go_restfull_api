package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// Book represents a book entity with ID, Title, Author, and Published Date.
type Book struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"publishedate"`
	Language      string `json:"language"`
}

var db *sqlx.DB

func init() {
	// Open a database connection
	var err error
	db, err = sqlx.Open("postgres", "user=postgres password=admin1234 host=127.0.0.1 dbname=book_rest_api sslmode=disable")
	if err != nil {
		panic(err)
	}

	// Check if the connection to the database is successful
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Print("Server is running on port(2222) \n")
	fmt.Println("Successfully! Connected to the database....")
}

func main() {
	// Create a NewRouter instance
	chi_router := chi.NewRouter()
	chi_router.Use(middleware.Logger)

	chi_router.Get("/getallbooks", getBooks)
	chi_router.Get("/getbooksbyid/{id}", getBookByID)
	chi_router.Post("/createbooks", createBook)
	chi_router.Delete("/deletebooksbyid/{id}", deleteBook)
	chi_router.Put("/updatebooksbyid/{id}", updateBook)

	http.ListenAndServe("localhost:2222", chi_router)

}

func createBook(w http.ResponseWriter, r *http.Request) {

	var newBook Book
	w.Header().Set("Content-Type", "application/json")
	// Parse the request body to create a new book
	json.NewDecoder(r.Body).Decode(&newBook)

	// Insert the new book into the database
	db.Exec("INSERT INTO books (id, title, author, publishedate) VALUES ($1, $2, $3, $4, $5)",
		newBook.ID, newBook.Title, newBook.Author, newBook.PublishedDate, newBook.Language)

	// Return the newly created book as JSON
	json.NewEncoder(w).Encode(newBook)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []Book

	value, err := db.Query("SELECT * FROM books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	defer value.Close()

	for value.Next() {
		var book Book
		value.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedDate, &book.Language)
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Placeholder for fetched data, replace with actual implementation
	var book Book
	// Retrieve a book by ID from the database and return it as JSON
	json.NewEncoder(w).Encode(book)

	// database query here
	err := db.QueryRow("SELECT * FROM books WHERE id = $1", chi.URLParam(r, "id")).
		Scan(&book.ID, &book.Title, &book.Author, &book.PublishedDate, &book.Language)
	// errors handler
	if err != nil {
		log.Fatal(err)
	}

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	// Delete a book by ID from the database
	_, err := db.Exec("DELETE FROM books WHERE id = $1", chi.URLParam(r, "id"))
	//  errors handler
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusNoContent)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var updatedBook Book
	w.Header().Set("Content-Type", "application/json")
	// Parse the request body to create a new book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Invalid Request Daata", http.StatusBadRequest)
		return
	}

	//database update query here
	db.Exec("UPDATE books SET title = $1, author = $2, publishedate = $3, language= $4 WHERE id = $5",
		updatedBook.Title, updatedBook.Author, updatedBook.PublishedDate, updatedBook.Language, chi.URLParam(r, "id"))

	// Update a book by ID in the database and return it as JSON
	json.NewEncoder(w).Encode(updatedBook)
	w.WriteHeader(http.StatusOK)

}
