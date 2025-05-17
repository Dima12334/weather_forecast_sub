package service

import (
	"context"
	"weather_forecast_sub/internal/config"
	"weather_forecast_sub/internal/repository"
	"weather_forecast_sub/pkg/clients"
	"weather_forecast_sub/pkg/email"
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

type ConfirmationEmailInput struct {
	Email string
	Token string
}

type WeatherForecastEmailInput struct {
}

type Emails interface {
	SendConfirmationEmail(ConfirmationEmailInput) error
	SendWeatherForecastEmail(WeatherForecastEmailInput) error
}

type Deps struct {
	Repos       *repository.Repositories
	Clients     *clients.Clients
	EmailHasher hash.EmailHasher
	EmailSender email.Sender
	EmailConfig config.EmailConfig
	HTTPConfig  config.HTTPConfig
}

type Services struct {
	Subscriptions Subscription
	Weather       Weather
}

func NewServices(deps Deps) *Services {
	emailsService := NewEmailsService(deps.EmailSender, deps.EmailConfig, deps.HTTPConfig)
	return &Services{
		Subscriptions: NewSubscriptionService(
			deps.Repos.Subscription,
			deps.EmailHasher,
			deps.EmailSender,
			deps.EmailConfig,
			deps.HTTPConfig,
			emailsService,
		),
		Weather: NewWeatherService(deps.Clients.WeatherAPI),
	}
}
