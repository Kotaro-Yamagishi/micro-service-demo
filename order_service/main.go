// order_service/main.go
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

type Order struct {
    ID     int    `json:"id"`
    UserID int    `json:"user_id"`
    BookID int    `json:"book_id"`
}

var orders = []Order{
    {ID: 1, UserID: 1, BookID: 1},
}

func getOrders(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(orders)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for _, order := range orders {
        if order.ID == id {
            json.NewEncoder(w).Encode(order)
            return
        }
    }
    http.NotFound(w, r)
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/orders", getOrders).Methods("GET")
    router.HandleFunc("/orders/{id}", getOrder).Methods("GET")
    log.Fatal(http.ListenAndServe(":8002", router))
}
