package handler

import (
	"net/http"
	"time"

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

	accessToken, refershToken, err := h.service.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, JSON{"error": err.Error()})
	}

	cookie := new(http.Cookie)
	cookie.Name = "refersh_token"
	cookie.Value = refershToken
	cookie.Expires = time.Now().Add(24 * 7 * time.Hour)
	cookie.Secure = false
	cookie.SameSite = http.SameSiteDefaultMode
	cookie.HttpOnly = true
	cookie.Path = "/refersh"

	c.SetCookie(cookie)

	return c.JSON(http.StatusCreated, JSON{"message": "Created Successfully", "token": accessToken})
}

func (h *Handler) SignIn(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, JSON{"error": err.Error()})
	}

	accessToken, refershToken, err := h.service.LoginUser(user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, JSON{"error": err.Error()})
	}

	cookie := new(http.Cookie)
	cookie.Name = "refresh_token"
	cookie.Value = refershToken
	cookie.Expires = time.Now().Add(24 * 7 * time.Hour)
	cookie.Secure = false
	cookie.SameSite = http.SameSiteDefaultMode
	cookie.HttpOnly = false
	cookie.Path = "/api/"

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, JSON{"message": "Logged in Successfully", "access_token": accessToken})
}

func (h *Handler) Refersh(c echo.Context) error {
	cookie, err := c.Cookie("refresh_token")

	if err != nil {
		return c.JSON(http.StatusUnauthorized, JSON{"error": err.Error()})
	}

	accessToken, err := h.service.RefershToken(cookie.Value)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, JSON{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, JSON{"message": "Logged in Successfully", "access_token": accessToken})

}
