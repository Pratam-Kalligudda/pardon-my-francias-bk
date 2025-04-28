package main

import (
	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/internal/config"
	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/internal/handler"
	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/internal/repo"
	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/internal/services"
	"github.com/labstack/echo/v4"
)

func main() {
	DB := config.InitDB()
	repo := repo.NewRepo(DB)
	service := services.NewService(&repo)
	handler := handler.NewHandler(&service)

	e := echo.New()
	e.POST("/api/signup", handler.SignUp)
	e.POST("/api/signin", handler.SignIn)
	e.Logger.Fatal(e.Start(":8080"))

}
