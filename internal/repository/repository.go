package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
	"weather_forecast_sub/internal/domain"
)

type SubscriptionRepository interface {
	Create(ctx context.Context, subscription domain.Subscription) error
	GetByToken(ctx context.Context, token string) (domain.Subscription, error)
	Confirm(ctx context.Context, token string) error
	SetLastSentAt(ctx context.Context, lastSentAt time.Time, token string) error
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
