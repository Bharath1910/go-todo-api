package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitiateRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/todos", getAllTodos).Methods("GET")
	r.HandleFunc("/todo/{username}", getTodo).Methods("GET")

	http.ListenAndServe(":8000", r)
}

func main() {

}
