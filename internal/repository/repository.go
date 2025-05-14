package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"weather_forecast_sub/internal/domain"
)

type UpdateSubscriptionInput struct {
	ID        string `json:"id"`
	Token     string `json:"token"`
	Confirmed bool   `json:"confirmed"`
}

type SubscriptionRepository interface {
	Create(ctx context.Context, subscription domain.Subscription) error
	GetByToken(ctx context.Context, token string) (domain.Subscription, error)
	Update(ctx context.Context, inp UpdateSubscriptionInput) error
	Delete(ctx context.Context, token string) error
}

type Repositories struct {
	Subscription SubscriptionRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Subscription: NewSubscriptionRepo(db),
	}
}
