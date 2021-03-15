package main

import (
	toDo "Routes/Todo"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", toDo.GetAllTodos)
	log.Fatal(http.ListenAndServe(":3333", nil))
}
