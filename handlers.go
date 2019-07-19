package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to go server!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		{Name: "Write presentation"},
		{Name: "Host meetup"},
	}
	json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	todoList := mux.Vars(r)
	todoId := todoList["todoId"]
	fmt.Fprintln(w, "Todo show: ", todoId)
}
