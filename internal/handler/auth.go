package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/models"
	"github.com/labstack/echo/v4"
)

type JSON map[string]any

func (h *Handler) SignUp(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(&user); err != nil {
		return err
	}

	accessToken, refershToken, err := h.service.CreateUser(user)

	jsonPtr, _ := json.MarshalIndent(user, "", "\t")
	fmt.Print(string(jsonPtr))

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

	return c.JSON(http.StatusOK, JSON{"message": "Logged in Successfully", "token": accessToken})
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

	return c.JSON(http.StatusOK, JSON{"message": "Logged in Successfully", "token": accessToken})

}

func (h *Handler) SignOut(c echo.Context) error {
	_, err := c.Cookie("refresh_token")

	if err != nil {
		return c.JSON(http.StatusBadRequest, JSON{"error": "No refresh token found : " + err.Error()})
	}

	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		MaxAge:   -1,
		Secure:   false,
		SameSite: http.SameSiteDefaultMode,
		Path:     "/api/",
		Expires:  time.Unix(0, 0),
		HttpOnly: false,
	})
	return c.JSON(http.StatusOK, JSON{"message": "Logged out successfuly"})
}
