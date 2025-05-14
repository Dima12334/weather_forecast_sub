package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) subscribeEmail(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "subscription successful",
	})
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
