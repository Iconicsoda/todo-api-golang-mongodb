package todo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/iconicsoda/todo-api-golang-mongodb/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type toDo struct {
	ID          string
	Title       string
	Description string
}

var data *mongo.Client = database.GetClient()

func (toDo *toDo) createID() {
	toDo.ID = uuid.NewString()
}

func GetAllTodos(resp http.ResponseWriter, req *http.Request) {
	toDos := returnAllToDos(data, bson.M{})

	resp.WriteHeader(http.StatusOK)

	json.NewEncoder(resp).Encode(toDos)
}

func GetOneTodo(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	todo := returnOneTodo(data, bson.M{"id": vars["id"]})
	resp.WriteHeader(http.StatusAccepted)
	json.NewEncoder(resp).Encode(todo)
}

func PostToDo(resp http.ResponseWriter, req *http.Request) {
	var newTodo toDo

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(reqBody, &newTodo)
	newTodo.createID()

	resp.WriteHeader(http.StatusCreated)

	insertNewTodo(data, newTodo)
	json.NewEncoder(resp).Encode(newTodo)
}

func returnAllToDos(client *mongo.Client, filter bson.M) []*toDo {

	var toDos []*toDo
	collection := client.Database("apiTodo").Collection("toDo")
	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}

	for cur.Next(context.TODO()) {
		var todo toDo
		err = cur.Decode(&todo)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		toDos = append(toDos, &todo)
	}

	return toDos
}

func returnOneTodo(client *mongo.Client, filter bson.M) toDo {
	var todo toDo
	collection := client.Database("apiTodo").Collection("toDo")
	result := collection.FindOne(context.TODO(), filter)
	result.Decode(&todo)
	return todo
}

func insertNewTodo(client *mongo.Client, todo toDo) interface{} {
	collection := client.Database("apiTodo").Collection("toDo")
	insertResult, err := collection.InsertOne(context.TODO(), todo)

	if err != nil {
		log.Fatal(err)
	}

	return insertResult.InsertedID
}
