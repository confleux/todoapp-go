package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"todoapp-go/internal/entities"
)

type FeedbackRepository struct {
	Pool *pgxpool.Pool
}

func NewFeedbackRepository(pool *pgxpool.Pool) *FeedbackRepository {
	return &FeedbackRepository{Pool: pool}
}

func (repo *FeedbackRepository) CreateFeedbackItem(ctx context.Context, email string, text string) (entities.FeedbackItem, error) {
	var todo entities.FeedbackItem

	if err := repo.Pool.QueryRow(ctx, "INSERT INTO public.feedback_item (email, text) VALUES ($1, $2) RETURNING email, text", email, text).Scan(&todo.Email, &todo.Text); err != nil {
		return entities.FeedbackItem{}, fmt.Errorf("unable to create todo_item: %w", err)
	}

	return todo, nil
}

func (repo *FeedbackRepository) GetAllFeedbackItems(ctx context.Context) ([]entities.FeedbackItem, error) {
	res := make([]entities.FeedbackItem, 0)

	rows, _ := repo.Pool.Query(ctx, "SELECT email, text FROM public.feedback_item")
	defer rows.Close()

	for rows.Next() {
		tmp := entities.FeedbackItem{}

		if err := rows.Scan(&tmp.Email, &tmp.Text); err != nil {
			return nil, fmt.Errorf("unable to get feedback_items: %w", err)
		}

		res = append(res, tmp)
	}

	return res, nil
}
