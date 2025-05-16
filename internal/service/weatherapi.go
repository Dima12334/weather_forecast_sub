package service

import (
	"context"
	"weather_forecast_sub/pkg/clients"
)

type WeatherService struct {
	client clients.Weather
}

func NewWeatherService(client clients.Weather) *WeatherService {
	return &WeatherService{client: client}
}

func (s *WeatherService) GetCurrentWeather(ctx context.Context, city string) (*clients.CurrentWeatherResponse, error) {
	return s.client.GetCurrentWeather(city)
}
