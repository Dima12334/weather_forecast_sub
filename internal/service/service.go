package service

import (
	"context"
	"weather_forecast_sub/internal/domain"
	"weather_forecast_sub/internal/repository"
)

type UpdateSubscriptionInput struct {
	ID        string `json:"id"`
	Token     string `json:"token"`
	Confirmed bool   `json:"confirmed"`
}

type Subscription interface {
	Create(ctx context.Context, subscription domain.Subscription) error
	Update(ctx context.Context, inp UpdateSubscriptionInput) error
	Delete(ctx context.Context, token string) error
}

type Deps struct {
	Repos *repository.Repositories
}

type Services struct {
	Subscriptions Subscription
}

func NewServices(deps Deps) *Services {
	return &Services{
		Subscriptions: NewSubscriptionService(deps.Repos.Subscription),
	}
}
