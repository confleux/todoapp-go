package repository

import (
	"context"
	"fmt"
	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type TodoRepository struct {
	Pool *pgxpool.Pool
}

func NewTodoRepository(pool *pgxpool.Pool) *TodoRepository {
	return &TodoRepository{Pool: pool}
}

func (repo *TodoRepository) CreateTodoItem(ctx context.Context, userId uuid.UUID, username string) (uuid.UUID, error) {
	createdAt := time.Now()

	var res uuid.UUID

	if err := repo.Pool.QueryRow(ctx, "INSERT INTO user (user_id, username, created_at) values (?, ?, ?) returning user_id", userId, username, createdAt).Scan(&res); err != nil {
		return uuid.UUID{}, fmt.Errorf("unable to get string: %w", err)
	}

	return userId, nil
}
