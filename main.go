package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
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

type TodoData struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Uuid      string `json:"uuid"`
}
type Users struct {
	Username string     `json:"username"`
	Password string     `json:"password"`
	TodoData []TodoData `json:"todoData"`
}

func getAllTodos(res http.ResponseWriter, req *http.Request) {
	allTodo := []Users{}

	collection := collection("Todo", "users")
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}

	for cur.Next(context.Background()) {
		var todo Users
		err := cur.Decode(&todo)
		if err != nil {
			fmt.Println(err)
		}
		allTodo = append(allTodo, todo)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	todoJSON, err := json.Marshal(allTodo)
	if err != nil {
		fmt.Println(err)
	}

	res.Write(todoJSON)

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
