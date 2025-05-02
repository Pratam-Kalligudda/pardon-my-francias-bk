package handler

import (
	"net/http"

	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) AddNote(c echo.Context) error {
	var note models.Note

	if err := c.Bind(&note); err != nil {
		return c.JSON(http.StatusBadRequest, JSON{"error": "no note provided : " + err.Error()})
	}

	tokenString := c.Request().Header.Get("Authorization")
	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, JSON{"error": "no token provided"})
	}

	tokenString = tokenString[len("Bearer "):]
	err := h.service.CreateNote(&note, tokenString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, JSON{"error": "something went wrong : " + err.Error()})
	}

	return c.JSON(http.StatusOK, JSON{"message": " note created succesfully", "note": note})
}

func (h *Handler) GetNotesOfUser(c echo.Context) error {
	tokenString := c.Request().Header.Get("Authorization")
	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, JSON{"error": "no token provided"})
	}

	tokenString = tokenString[len("Bearer "):]
	notes, err := h.service.GetNotesOfUser(tokenString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, JSON{"error": "something went wrong : " + err.Error()})
	}

	return c.JSON(http.StatusOK, JSON{"message": "success", "notes": notes})
}
