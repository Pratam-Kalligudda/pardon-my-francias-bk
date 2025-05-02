package handler

import "github.com/Pratam-Kalligudda/pardon-my-francias-bk/internal/services"

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) Handler {
	return Handler{service}
}
