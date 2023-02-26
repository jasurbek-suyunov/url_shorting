package handler

import "github.com/SuyunovJasurbek/url_shorting/src/service"

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
