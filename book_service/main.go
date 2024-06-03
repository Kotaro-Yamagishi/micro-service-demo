// book_service/main.go
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"

    `github.com/gorilla/mux`
)

type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var books = []Book{
    {ID: 1, Title: "The Go Programming Language", Author: "Alan A. A. Donovan"},
    {ID: 2, Title: "Learning Go", Author: "Jon Bodner"},
}

func getBooks(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for _, book := range books {
        if book.ID == id {
            json.NewEncoder(w).Encode(book)
            return
        }
    }
    http.NotFound(w, r)
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/books", getBooks).Methods("GET")
    router.HandleFunc("/books/{id}", getBook).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}
