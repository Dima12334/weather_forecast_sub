package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"weather_forecast_sub/internal/domain"
	customErrors "weather_forecast_sub/pkg/errors"
)

type SubscriptionRepo struct {
	db *sqlx.DB
}

func NewSubscriptionRepo(db *sqlx.DB) *SubscriptionRepo {
	return &SubscriptionRepo{db: db}
}

func (r *SubscriptionRepo) Create(ctx context.Context, subscription domain.Subscription) error {
	query := `
		INSERT INTO subscriptions (created_at, email, city, token, frequency, confirmed, last_sent_at) 
		values ($1, $2, $3, $4, $5, $6, $7);`
	_, err := r.db.QueryxContext(
		ctx,
		query,
		subscription.CreatedAt,
		subscription.Email,
		subscription.City,
		subscription.Token,
		subscription.Frequency,
		subscription.Confirmed,
		subscription.LastSentAt,
	)
	if err != nil {
		if customErrors.IsDuplicateDBError(err) {
			return customErrors.ErrSubscriptionAlreadyExists
		}
	}
	return err
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
