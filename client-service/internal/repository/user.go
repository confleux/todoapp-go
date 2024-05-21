package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	Pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{Pool: pool}
}

func (repo *UserRepository) CreateUser(ctx context.Context, uid string, email string) (string, error) {
	var res string

	if err := repo.Pool.QueryRow(ctx, "INSERT INTO public.user (uid, email) VALUES ($1, $2) RETURNING uid", uid, email).Scan(&res); err != nil {
		return "", fmt.Errorf("unable to create user: %w", err)
	}

	return res, nil
}
