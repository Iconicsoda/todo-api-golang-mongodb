package todo

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type toDo struct {
	ID          string
	Title       string
	Description string
}

type allToDos []toDo

var toDos = allToDos{
	{
		ID:          uuid.NewString(),
		Title:       "Test",
		Description: "Test Description",
	},
}

func (toDo *toDo) createID() {
	toDo.ID = uuid.NewString()
}

func GetAllTodos(resp http.ResponseWriter, req *http.Request) {
	json.NewEncoder(resp).Encode(toDos)
}
