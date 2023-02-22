package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func collection(database string, coll string) *mongo.Collection {
	fmt.Println("Connecting to MongoDB...")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}
	collection := client.Database(database).Collection(coll)
	fmt.Println("Connected to MongoDB!")

	return collection
}

func getAllTodos(res http.ResponseWriter, req *http.Request) {

}

func getTodo(res http.ResponseWriter, req *http.Request) {

}

func InitiateRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/todos", getAllTodos).Methods("GET")
	r.HandleFunc("/todo/{username}", getTodo).Methods("GET")

	http.ListenAndServe(":8000", r)
}

func main() {
	InitiateRouter()
}
