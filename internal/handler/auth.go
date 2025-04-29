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

	token, err := h.service.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, JSON{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, JSON{"message": "Created Successfully", "token": token})
}

func (h *Handler) SignIn(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, JSON{"error": err.Error()})
	}

	token, err := h.service.LoginUser(user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, JSON{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, JSON{"message": "Logged in Successfully", "token": token})
}
