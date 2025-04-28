package handler

import (
	"net/http"

	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/internal/services"
	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/models"
	"github.com/labstack/echo/v4"
)

type JSON map[string]string

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) Handler {
	return Handler{service}
}

func (h *Handler) SignUp(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := h.service.CreateUser(user); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, JSON{"message": "Created Successfully"})
}

func (h *Handler) SignIn(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := h.service.LoginUser(user.Email, user.Password); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, JSON{"message": "Logged in Successfully"})
}
