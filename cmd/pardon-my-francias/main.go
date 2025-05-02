package main

import (
	"fmt"
	"os"

	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/internal/config"
	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/internal/handler"
	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/internal/repo"
	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/internal/services"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	port := ":" + os.Getenv("PORT")
	DB := config.InitDB()
	repo := repo.NewRepo(DB)
	service := services.NewService(&repo)
	handler := handler.NewHandler(&service)

	e := echo.New()
	e.POST("/api/signup", handler.SignUp)
	e.POST("/api/signin", handler.SignIn)
	e.POST("/api/refresh", handler.Refersh)
	e.POST("/api/signout", handler.SignOut)
	e.Logger.Fatal(e.Start(port))

}
