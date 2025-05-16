package service

import (
	"context"
	"weather_forecast_sub/internal/repository"
	"weather_forecast_sub/pkg/clients"
	"weather_forecast_sub/pkg/hash"
)

type CreateSubscriptionInput struct {
	Email     string `json:"email"`
	City      string `json:"city"`
	Frequency string `json:"frequency"`
}

type Subscription interface {
	Create(ctx context.Context, inp CreateSubscriptionInput) error
	Confirm(ctx context.Context, token string) error
	Delete(ctx context.Context, token string) error
}

type Weather interface {
	GetCurrentWeather(ctx context.Context, city string) (*clients.CurrentWeatherResponse, error)
}

type Deps struct {
	Repos   *repository.Repositories
	Clients *clients.Clients
	Hasher  hash.EmailHasher
}

type Services struct {
	Subscriptions Subscription
	Weather       Weather
}

func NewServices(deps Deps) *Services {
	return &Services{
		Subscriptions: NewSubscriptionService(deps.Repos.Subscription, deps.Hasher),
		Weather:       NewWeatherService(deps.Clients.WeatherAPI),
	}
}
