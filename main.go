package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var (
	books      = make(map[int]Book)
	nextID     = 1
	booksMutex = &sync.Mutex{}
)

func loadSampleBooks() {
	samples := []Book{
		{Title: "1984", Author: "George Orwell"},
		{Title: "To Kill a Mockingbird", Author: "Harper Lee"},
		{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
	}

	booksMutex.Lock()
	defer booksMutex.Unlock()

	for _, b := range samples {
		b.ID = nextID
		books[nextID] = b
		nextID++
	}
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	booksMutex.Lock()
	defer booksMutex.Unlock()

	var all []Book
	for _, book := range books {
		all = append(all, book)
	}
	json.NewEncoder(w).Encode(all)
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	booksMutex.Lock()
	defer booksMutex.Unlock()

	book, exists := books[id]
	if !exists {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	booksMutex.Lock()
	defer booksMutex.Unlock()

	book.ID = nextID
	nextID++
	books[book.ID] = book

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	var updated Book
	json.NewDecoder(r.Body).Decode(&updated)

	booksMutex.Lock()
	defer booksMutex.Unlock()

	if _, exists := books[id]; !exists {
		http.NotFound(w, r)
		return
	}
	updated.ID = id
	books[id] = updated
	json.NewEncoder(w).Encode(updated)
}

func main() {
	loadSampleBooks()

	router := mux.NewRouter()
	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBookByID).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")

	log.Println("Server running on :7070")
	log.Fatal(http.ListenAndServe(":7070", router))
}
