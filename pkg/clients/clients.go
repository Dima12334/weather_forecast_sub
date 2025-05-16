package clients

import (
	"weather_forecast_sub/internal/config"
)

type Weather interface {
	GetCurrentWeather(city string) (*CurrentWeatherResponse, error)
}

type Clients struct {
	WeatherAPI Weather
}

func NewClients(thirdPartyCfg config.ThirdPartyConfig) *Clients {
	return &Clients{
		WeatherAPI: NewWeatherAPIClient(thirdPartyCfg.WeatherAPIKey),
	}
}
