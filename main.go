package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitiateRouter() {
	r := mux.NewRouter()
	r.HandleFunc()
	r.HandleFunc()

	http.ListenAndServe(":8000", r)
}

func main() {

}
