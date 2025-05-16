package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	customErrors "weather_forecast_sub/pkg/errors"
)

func (h *Handler) getWeather(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	escapedCity := url.QueryEscape(city)

	weather, err := h.services.Weather.GetCurrentWeather(c.Request.Context(), escapedCity)
	if err != nil {
		switch {
		case errors.Is(err, customErrors.ErrCityNotFound):
			c.Status(http.StatusNotFound)
		default:
			c.Status(http.StatusBadRequest)
		}
		return
	}

	c.JSON(http.StatusOK, weather)
}
