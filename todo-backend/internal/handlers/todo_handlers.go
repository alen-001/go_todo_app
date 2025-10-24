package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"todo-backend/internal/database"
	"todo-backend/internal/models"
)

func TodoHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type","application/json")
	switch req.Method {
	case "POST":
		addTodo(w, req)
	case "GET":
		if parts := strings.Split(req.URL.Path, "/");parts[2]==""{
			getTodos(w, req)
		}else{
			id:=parts[2]
			getTodoById(w,req,id)
		}
	case "PUT":
		if parts := strings.Split(req.URL.Path, "/");parts[2]==""{
			http.Error(w,"Id is required",http.StatusBadRequest)
		}
	case "DELETE":
		if parts :=strings.Split(req.URL.Path,"/");parts[2]=="all"{
			deleteAllTodos(w,req)
		}else if parts[2]==""{
			http.Error(w,"A path paramater must be passed",http.StatusBadRequest)
		}else{
			id:=parts[2]
			deleteTodoById(w,req,id);
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func getTodos(w http.ResponseWriter, req *http.Request) {
	todos:=database.Todos
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
func addTodo(w http.ResponseWriter, req *http.Request) {
	var payload models.Todo
	if err:=json.NewDecoder(req.Body).Decode(&payload);err!=nil{
		http.Error(w,"Invalid request payload",http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	sz := len(database.Todos)
	database.Todos[sz+1]=models.Todo{Id:payload.Id,Name:payload.Name,Status:payload.Status}
	fmt.Fprintf(w, "Todo added successfully: %+v\n", payload)
}

func getTodoById(w http.ResponseWriter, req *http.Request,id string) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	value,ok:=database.Todos[idInt]
	if !ok{
		http.Error(w,"Todo not found",http.StatusNotFound)
		return
	}
	res:= map[string]interface{}{
		"Message":"Todo found successfully",
		"Todo":value,
	}
	json.NewEncoder(w).Encode(res)
}
func deleteAllTodos(w http.ResponseWriter, req *http.Request){

}
func deleteTodoById(w http.ResponseWriter,req * http.Request,id string){}