package controller

import (
	"client-service/internal/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"client-service/internal/entities"
)

type TodoController struct {
	todoService *service.TodoService
}

func NewTodoController(todoService *service.TodoService) *TodoController {
	return &TodoController{todoService: todoService}
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

// CreateTodo godoc
//
//	@Summary Create todo item
//	@Description	Create todo item
//	@Tags         todo
//	@Accept       json
//	@Produce      json
//	@Param todo body CreateTodoRequest true "Todo description"
//	@Success		200	{object}	entities.Todo
//	@Failure		400	{object}	nil "Bad Request"
//	@Failure		401	{object}	nil "Unauthorized"
//	@Failure		500	{object}	nil "Internal Server Error"
//	@Router /todos [post]
//
// @Security Bearer
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

	todo, err := ct.todoService.CreateTodoItem(r.Context(), req.Description, uid)
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

// GetTodo godoc
//
//	@Summary Get todo item
//	@Description	Get todo item
//	@Tags         todo
//	@Accept       json
//	@Produce      json
//	@Param id path string true "Todo uuid"
//	@Success		200	{object}	entities.Todo
//	@Failure		400	{object}	nil "Bad Request"
//	@Failure		401	{object}	nil "Unauthorized"
//	@Failure		500	{object}	nil "Internal Server Error"
//	@Router /todos/{id} [get]
//
// @Security Bearer
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

	todo, err := ct.todoService.GetTodo(r.Context(), todoId)
	if err != nil {
		http.Error(w, "Unable to get todo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(&todo); err != nil {
		http.Error(w, "Unable to get todo", http.StatusInternalServerError)
		return
	}
}

// GetTodos godoc
//
//	@Summary Get todo items
//	@Description	Get todo items associated with user
//	@Tags         todo
//	@Accept       json
//	@Produce      json
//	@Success		200	{object} GetTodosResponse
//	@Failure		400	{object}	nil "Bad Request"
//	@Failure		401	{object}	nil "Unauthorized"
//	@Failure		500	{object}	nil "Internal Server Error"
//	@Router /todos [get]
//
// @Security Bearer
func (ct *TodoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	var res GetTodosResponse

	uid := r.Context().Value("uid").(string)

	todos, err := ct.todoService.GetTodos(r.Context(), uid)
	if err != nil {
		http.Error(w, "Unable to get todos", http.StatusInternalServerError)
		return
	}

	res.Todos = todos

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(&res); err != nil {
		http.Error(w, "Unable to get todos", http.StatusInternalServerError)
		return
	}
}

// RemoveTodo godoc
//
//	@Summary Remove todo item
//	@Description	Remove todo item by id
//	@Tags         todo
//	@Accept       json
//	@Produce      json
//	@Param id path string true "Todo uuid"
//	@Success		200	{object} entities.Todo
//	@Failure		400	{object}	nil "Bad Request"
//	@Failure		401	{object}	nil "Unauthorized"
//	@Failure		500	{object}	nil "Internal Server Error"
//	@Router /todos/{id} [delete]
//
// @Security Bearer
func (ct *TodoController) RemoveTodo(w http.ResponseWriter, r *http.Request) {
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

	todo, err := ct.todoService.RemoveTodo(r.Context(), todoId)
	if err != nil {
		http.Error(w, "Unable to remove todo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(&todo); err != nil {
		http.Error(w, "Unable to remove todo", http.StatusInternalServerError)
		return
	}
}
