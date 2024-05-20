package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	Pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{Pool: pool}
}

func (repo *UserRepository) CreateUser(ctx context.Context, uid string, email string) (string, error) {
	createdAt := time.Now()

	var res string

	if err := repo.Pool.QueryRow(ctx, "INSERT INTO public.user (uid, email, created_at) VALUES ($1, $2, $3) RETURNING uid", uid, email, createdAt).Scan(&res); err != nil {
		return "", fmt.Errorf("unable to create user: %w", err)
	}

	fmt.Println(res)

	return res, nil
}
