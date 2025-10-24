package main

import (
	"log"
	"net/http"
	h "todo-backend/internal/handlers"
)
func main() {
	http.HandleFunc("/todo/", h.TodoHandler)
	log.Println("Server started on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
