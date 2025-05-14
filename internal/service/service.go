package service

import (
	"context"
	"weather_forecast_sub/internal/repository"
	"weather_forecast_sub/pkg/hash"
)

type CreateSubscriptionInput struct {
	Email     string `json:"email"`
	City      string `json:"city"`
	Frequency string `json:"frequency"`
}

type UpdateSubscriptionInput struct {
	ID        string `json:"id"`
	Token     string `json:"token"`
	Confirmed bool   `json:"confirmed"`
}

type Subscription interface {
	Create(ctx context.Context, inp CreateSubscriptionInput) error
	Update(ctx context.Context, inp UpdateSubscriptionInput) error
	Delete(ctx context.Context, token string) error
}

type Deps struct {
	Repos  *repository.Repositories
	Hasher hash.EmailHasher
}

type Services struct {
	Subscriptions Subscription
}

func NewServices(deps Deps) *Services {
	return &Services{
		Subscriptions: NewSubscriptionService(deps.Repos.Subscription, deps.Hasher),
	}
}
