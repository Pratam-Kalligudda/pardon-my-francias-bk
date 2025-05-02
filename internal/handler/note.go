package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateNote(c echo.Context) error {
	_, ok := c.Request().Header["Authorization"]
	if !ok {
		return c.JSON(http.StatusBadRequest, JSON{"error": "No token provided : "})
	}

	// claims, err := h.service.

	return c.JSON(http.StatusOK, JSON{"message": " note created succesfully"})
}
