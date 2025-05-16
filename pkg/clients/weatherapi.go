package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	customErrors "weather_forecast_sub/pkg/errors"
	"weather_forecast_sub/pkg/logger"
)

type WeatherAPIClient struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

func NewWeatherAPIClient(apiKey string) *WeatherAPIClient {
	return &WeatherAPIClient{
		APIKey:     apiKey,
		BaseURL:    "https://api.weatherapi.com/v1",
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

type weatherAPIErrorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type CurrentWeatherResponse struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Description string  `json:"description"`
}

type currentWeatherAPIResponse struct {
	Current struct {
		TempC     float32 `json:"temp_c"`
		Humidity  float32 `json:"humidity"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

func (c *WeatherAPIClient) GetCurrentWeather(city string) (*CurrentWeatherResponse, error) {
	url := fmt.Sprintf("%s/current.json?key=%s&q=%s", c.BaseURL, c.APIKey, city)

	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		logger.Errorf("Error making request to Weather API: %s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErr weatherAPIErrorResponse
		if err = json.NewDecoder(resp.Body).Decode(&apiErr); err == nil {
			if resp.StatusCode == http.StatusBadRequest && apiErr.Error.Code == 1006 {
				return nil, customErrors.ErrCityNotFound
			}
			logger.Errorf(
				"Weather API error. Status code: %d, api code %d, message %s",
				resp.StatusCode,
				apiErr.Error.Code,
				apiErr.Error.Message,
			)
			return nil, customErrors.ErrWeatherAPIError
		}
		logger.Errorf("Weather API error. Status code: %d", resp.StatusCode)
		return nil, customErrors.ErrWeatherAPIError
	}

	var result currentWeatherAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		logger.Errorf("Error parsing Weather API response: %s", err.Error())
		return nil, customErrors.ErrWeatherAPIError
	}

	return &CurrentWeatherResponse{
		Temperature: result.Current.TempC,
		Humidity:    result.Current.Humidity,
		Description: result.Current.Condition.Text,
	}, nil
}
