package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"client-service/internal/entities"
)

type TodoRepository struct {
	Pool *pgxpool.Pool
}

func NewTodoRepository(pool *pgxpool.Pool) *TodoRepository {
	return &TodoRepository{Pool: pool}
}

func (repo *TodoRepository) CreateTodoItem(ctx context.Context, description string, uid string) (entities.Todo, error) {
	var todo entities.Todo

	id := uuid.New()

	if err := repo.Pool.QueryRow(ctx, "INSERT INTO public.todo_item (description, id, uid) VALUES ($1, $2, $3) RETURNING description, id, created_at, uid", description, id, uid).Scan(&todo.Description, &todo.Id, &todo.CreatedAt, &todo.Uid); err != nil {
		return entities.Todo{}, fmt.Errorf("unable to create todo_item: %w", err)
	}

	return todo, nil
}

func (repo *TodoRepository) GetTodoItemById(ctx context.Context, id uuid.UUID) (entities.Todo, error) {
	res := entities.Todo{}

	if err := repo.Pool.QueryRow(ctx, "SELECT description, id, created_at, uid  FROM public.todo_item WHERE id=$1", id).Scan(&res.Description, &res.Id, &res.CreatedAt, &res.Uid); err != nil {
		return entities.Todo{}, fmt.Errorf("unable to get todo_item by id: %w", err)
	}

	return res, nil
}

func (repo *TodoRepository) GetTodoItemsByUid(ctx context.Context, uid string) ([]entities.Todo, error) {
	res := make([]entities.Todo, 0)

	rows, _ := repo.Pool.Query(ctx, "SELECT description, id, created_at, uid  FROM public.todo_item WHERE uid=$1", uid)
	defer rows.Close()

	for rows.Next() {
		tmp := entities.Todo{}

		if err := rows.Scan(&tmp.Description, &tmp.Id, &tmp.CreatedAt, &tmp.Uid); err != nil {
			return nil, fmt.Errorf("unable to get todo_items by uid: %w", err)
		}

		res = append(res, tmp)
	}

	return res, nil
}
