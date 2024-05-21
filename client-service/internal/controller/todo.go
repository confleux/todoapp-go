package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"client-service/internal/entities"
	"client-service/internal/repository"
)

type TodoController struct {
	todoDB *repository.TodoRepository
}

func NewTodoController(userRepo *repository.TodoRepository) *TodoController {
	return &TodoController{todoDB: userRepo}
}

type CreateTodoRequest struct {
	Description string `json:"description"`
}

type CreateTodoResponse struct {
	Id string `json:"id"`
}

type GetTodosResponse struct {
	Todos []entities.Todo `json:"todos"`
}

func (ct *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req CreateTodoRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if len(req.Description) == 0 {
		http.Error(w, "Description cannot be empty", http.StatusBadRequest)
		return
	}

	uid := r.Context().Value("uid").(string)

	todo, err := ct.todoDB.CreateTodoItem(r.Context(), req.Description, uid)
	if err != nil {
		http.Error(w, "Unable to create todo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err = json.NewEncoder(w).Encode(&todo); err != nil {
		http.Error(w, "Unable to create todo", http.StatusInternalServerError)
		return
	}
}

func (ct *TodoController) GetTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	todoId, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	todo, err := ct.todoDB.GetTodoItemById(r.Context(), todoId)
	if err != nil {
		http.Error(w, "Unable to get todo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err = json.NewEncoder(w).Encode(&todo); err != nil {
		http.Error(w, "Unable to get todo", http.StatusInternalServerError)
		return
	}
}

func (ct *TodoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	var res GetTodosResponse

	uid := r.Context().Value("uid").(string)

	todos, err := ct.todoDB.GetTodoItemsByUid(r.Context(), uid)
	if err != nil {
		http.Error(w, "Unable to get todos", http.StatusInternalServerError)
		return
	}

	res.Todos = todos

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err = json.NewEncoder(w).Encode(&res); err != nil {
		http.Error(w, "Unable to get todos", http.StatusInternalServerError)
		return
	}
}