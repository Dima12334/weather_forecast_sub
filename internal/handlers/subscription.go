package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"weather_forecast_sub/internal/service"
	customErrors "weather_forecast_sub/pkg/errors"
)

type subscribeEmailInput struct {
	Email     string `json:"email" binding:"required,max=255"`
	City      string `json:"city" binding:"required,max=255"`
	Frequency string `json:"frequency" binding:"oneof=hourly daily"`
}

func (h *Handler) subscribeEmail(c *gin.Context) {
	var inp subscribeEmailInput

	if err := c.BindJSON(&inp); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err := h.services.Subscriptions.Create(
		c,
		service.CreateSubscriptionInput{
			Email:     inp.Email,
			City:      inp.City,
			Frequency: inp.Frequency,
		},
	)
	if err != nil {
		if errors.Is(err, customErrors.ErrSubscriptionAlreadyExists) {
			c.Status(http.StatusConflict)
			return
		}

		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) confirmEmail(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "subscription confirmed",
	})
}

func (h *Handler) unsubscribeEmail(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "subscription deleted",
	})
}
