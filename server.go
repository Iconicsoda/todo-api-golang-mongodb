package main

import (
	"log"
	"net/http"

	todo "github.com/iconicsoda/todo-api-golang-mongodb/Routes/Todo"
)

func main() {
	http.HandleFunc("/", todo.GetAllTodos)
	log.Fatal(http.ListenAndServe(":3333", nil))
}
