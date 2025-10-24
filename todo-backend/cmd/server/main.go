package main

import (
	"log"
	"fmt"
	"net/http"
	"strings"
	
)

// var todos []database.Todo

func getTodos(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "All todos requested!")
}

func addTodo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Add todo called")
}

func getTodoById(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	fmt.Println(parts)
	if parts[2] == "" {
		http.Error(w, "ID cannot be empty", http.StatusBadRequest)
		return
	}

	id := parts[2]
	fmt.Fprintf(w, "Todo with id: %s requested\n", id)
}

func todoHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		addTodo(w, req)
	case "GET":
		getTodoById(w, req)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func todosHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		getTodos(w, req)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	// t1 := database.Todo{Name: "bath"}
	// todos = append(todos, database.Todo{Id: 1, Name: "shower", Completed: true}, t1)

	http.HandleFunc("/todo/", todoHandler)
	http.HandleFunc("/todos", todosHandler)

	log.Println("Server started on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
