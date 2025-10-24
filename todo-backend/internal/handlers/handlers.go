package handlers
import(
	"fmt"
	"net/http"
	"strings"
)
func getTodos(w http.ResponseWriter, req *http.Request) {

}
// apis to be built 
// GET /todo --> Get all todos
//
func addTodo(w http.ResponseWriter, req *http.Request) {
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
