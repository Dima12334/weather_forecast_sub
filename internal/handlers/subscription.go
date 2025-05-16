package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"weather_forecast_sub/internal/service"
	customErrors "weather_forecast_sub/pkg/errors"
)

type subscribeEmailInput struct {
	Email     string `form:"email" json:"email" binding:"required,email,max=255"`
	City      string `form:"city" json:"city" binding:"required,max=255"`
	Frequency string `form:"frequency" json:"frequency" binding:"oneof=hourly daily"`
}

func (h *Handler) showSubscribePage(c *gin.Context) {
	c.HTML(http.StatusOK, "subscribe.html", gin.H{})
}

func (h *Handler) subscribeEmail(c *gin.Context) {
	var inp subscribeEmailInput

	if err := c.ShouldBind(&inp); err != nil {
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
	token := c.Param("token")

	err := h.services.Subscriptions.Confirm(c, token)
	if err != nil {
		switch {
		case errors.Is(err, customErrors.ErrSubscriptionNotFound):
			c.Status(http.StatusNotFound)
		default:
			c.Status(http.StatusBadRequest)
		}
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) unsubscribeEmail(c *gin.Context) {
	token := c.Param("token")

	err := h.services.Subscriptions.Delete(c, token)
	if err != nil {
		switch {
		case errors.Is(err, customErrors.ErrSubscriptionNotFound):
			c.Status(http.StatusNotFound)
		default:
			c.Status(http.StatusBadRequest)
		}
		return
	}

	c.Status(http.StatusOK)
}
