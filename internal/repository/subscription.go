package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"weather_forecast_sub/internal/domain"
)

type SubscriptionRepo struct {
	db *sqlx.DB
}

func NewSubscriptionRepo(db *sqlx.DB) *SubscriptionRepo {
	return &SubscriptionRepo{db: db}
}

func (r *SubscriptionRepo) Create(ctx context.Context, subscription domain.Subscription) error {
	return nil
}

func (r *SubscriptionRepo) GetByToken(ctx context.Context, token string) (domain.Subscription, error) {
	return domain.Subscription{}, nil
}

func (r *SubscriptionRepo) Update(ctx context.Context, inp UpdateSubscriptionInput) error {
	return nil
}

func (r *SubscriptionRepo) Delete(ctx context.Context, token string) error {
	return nil
}
