package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	todo "github.com/iconicsoda/todo-api-golang-mongodb/Routes/Todo"
	database "github.com/iconicsoda/todo-api-golang-mongodb/database"
)

func main() {

	route := mux.NewRouter().StrictSlash(true)

	route.HandleFunc("/todo/{id}", todo.GetOneTodo).Methods("GET")
	route.HandleFunc("/todo", todo.GetAllTodos).Methods("GET")
	route.HandleFunc("/todo", todo.PostToDo).Methods("POST")

	http.Handle("/", route)
	database.Database()
	log.Fatal(http.ListenAndServe(":3333", nil))
}
