package service

import (
	"client-service/internal/entities"
	"client-service/internal/repository"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type TodoService struct {
	todoDB *repository.TodoRepository
}

func NewTodoService(todoDB *repository.TodoRepository) *TodoService {
	return &TodoService{todoDB: todoDB}
}

func (s *TodoService) CreateTodoItem(ctx context.Context, description string, uid string) (entities.Todo, error) {
	todo, err := s.todoDB.CreateTodoItem(ctx, description, uid)
	if err != nil {
		return entities.Todo{}, fmt.Errorf("failed to create todo: %w", err)
	}
	return todo, nil
}

func (s *TodoService) GetTodo(ctx context.Context, todoId uuid.UUID) (entities.Todo, error) {
	todo, err := s.todoDB.GetTodoItemById(ctx, todoId)
	if err != nil {
		return entities.Todo{}, fmt.Errorf("failed to get todo: %w", err)
	}
	return todo, nil
}

func (s *TodoService) GetTodos(ctx context.Context, uid string) ([]entities.Todo, error) {
	todos, err := s.todoDB.GetTodoItemsByUid(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("failed to get todo: %w", err)
	}
	return todos, nil
}

func (s *TodoService) RemoveTodo(ctx context.Context, todoId uuid.UUID) (entities.Todo, error) {
	todo, err := s.todoDB.RemoveTodoItemById(ctx, todoId)
	if err != nil {
		return entities.Todo{}, fmt.Errorf("failed to remove todo: %w", err)
	}
	return todo, nil
}
