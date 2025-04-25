package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Todo represents a todo item
// swagger:model Todo
type Todo struct {
	// The ID of the todo
	// required: true
	ID int `json:"id"`
	// The title of the todo
	// required: true
	Title string `json:"title"`
	// Whether the todo is completed
	// required: true
	Completed bool `json:"completed"`
}

var (
	todos = map[int]Todo{
		1: {ID: 1, Title: "Learn Go", Completed: false},
		2: {ID: 2, Title: "Build REST API", Completed: false},
		3: {ID: 3, Title: "Write tests", Completed: false},
		4: {ID: 4, Title: "Document code", Completed: true},
		5: {ID: 5, Title: "Deploy application", Completed: false},
	}
	mu sync.RWMutex
	id = 6 // Start after the sample data
)

// swagger:route POST /todos todos createTodo
// Creates a new todo item.
// responses:
//
//	200: Todo
func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	todo.ID = id
	todos[id] = todo
	id++
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// swagger:route GET /todos todos listTodos
// Lists all todo items.
// responses:
//
//	200: []Todo
func listTodosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	mu.RLock()
	todoList := make([]Todo, 0, len(todos))
	for _, todo := range todos {
		todoList = append(todoList, todo)
	}
	mu.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todoList)
}

func main() {
	// Register handlers
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			listTodosHandler(w, r)
		case http.MethodPost:
			createTodoHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Starting server on port 8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("Server error:", err)
	}
}
