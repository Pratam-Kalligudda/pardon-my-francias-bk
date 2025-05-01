package services

import (
	"fmt"
	"time"

	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/internal/repo"
	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/models"
)

type Service struct {
	Repo *repo.Repo
}

func NewService(repo *repo.Repo) Service {
	return Service{repo}
}

var secrete string = "secret key"
var accessTokenDuration time.Duration = 24 * time.Hour
var refershTokenDuration time.Duration = 7 * 24 * time.Hour

func (s Service) CreateUser(user *models.User) (string, string, error) {
	users := s.Repo.GetUser("email", user.Email)
	if len(users) != 0 {
		return "", "", fmt.Errorf("email already exists")
	}
	users = nil
	users = s.Repo.GetUser("user_name", user.UserName)
	if len(users) != 0 {
		return "", "", fmt.Errorf("username already exists")
	}
	user.UserId = GenerateUserID()
	hashPass, err := GenerateHashPassword(user.Password)
	if err != nil {
		return "", "", err
	}
	user.Password = hashPass

	s.Repo.AddUser(*user)

	accessToken, err := GenerateJWTToken(user.UserId, accessTokenDuration)
	if err != nil {
		return "", "", err
	}

	refershToken, err := GenerateJWTToken(user.UserId, refershTokenDuration)
	if err != nil {
		return "", "", err
	}

	return accessToken, refershToken, nil
}

func (s Service) LoginUser(email, password string) (string, string, error) {
	users := s.Repo.GetUser("email", email)
	if len(users) > 1 || len(users) == 0 {
		return "", "", fmt.Errorf("email issue")
	}
	if err := ComparePassword(users[0].Password, password); err != nil {
		return "", "", err
	}

	accessToken, err := GenerateJWTToken(users[0].UserId, accessTokenDuration)
	if err != nil {
		return "", "", err
	}

	refershToken, err := GenerateJWTToken(users[0].UserId, refershTokenDuration)
	if err != nil {
		return "", "", err
	}
	return accessToken, refershToken, nil
}

func (s *Service) RefershToken(token string) (accessToken string, err error) {
	claims, err := ValidateJWTToken(token)
	if err != nil {
		return "", err
	}

	userId, err := claims.GetSubject()
	if err != nil {
		return "", err

	}

	accessToken, err = GenerateJWTToken(userId, accessTokenDuration)
	if err != nil {
		return "", err
	}

	return
}
