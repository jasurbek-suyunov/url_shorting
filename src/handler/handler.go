package handler

import (
	"github.com/SuyunovJasurbek/url_shorting/src/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (h *Handler) API(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "API version 1",
	})
}
