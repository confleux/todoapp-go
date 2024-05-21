package routes

import (
	"github.com/go-chi/chi/v5"

	"client-service/internal/controller"
)

type TodosResource struct {
	todoController *controller.TodoController
}

func NewTodoResource(todoController *controller.TodoController) *TodosResource {
	return &TodosResource{todoController: todoController}
}

// Routes creates a REST router for the todos resource
func (rs TodosResource) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", rs.todoController.GetTodos)    // GET /todos - read a list of todos
	r.Post("/", rs.todoController.CreateTodo) // POST /todos - create a new todo and persist it
	//r.Put("/", rs.Delete)

	r.Route("/{id}", func(r chi.Router) {
		// r.Use(rs.TodoCtx) // lets have a todos map, and lets actually load/manipulate
		r.Get("/", rs.todoController.GetTodo) // GET /todos/{id} - read a single todo by :id
		//r.Put("/", rs.Update)    // PUT /todos/{id} - update a single todo by :id
		r.Delete("/", rs.todoController.RemoveTodo) // DELETE /todos/{id} - delete a single todo by :id
		//r.Get("/sync", rs.Sync)
	})

	return r
}
