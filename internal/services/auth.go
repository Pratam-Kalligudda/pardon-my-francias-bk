package services

import (
	"fmt"
	"time"

	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/models"
)

var accessTokenDuration time.Duration = 24 * time.Hour
var refershTokenDuration time.Duration = 7 * 24 * time.Hour

func (s *Service) CreateUser(user *models.User) (string, string, error) {
	err := s.Repo.CheckIfUserExists("email", user.Email)
	if err == nil {
		return "", "", fmt.Errorf("email already exists : %v ", err)
	}
	err = s.Repo.CheckIfUserExists("user_name", user.UserName)
	if err == nil {
		return "", "", fmt.Errorf("username already exists : %v ", err)
	}

	hashPass, err := generateHashPassword(user.Password)
	if err != nil {
		return "", "", err
	}
	user.Password = hashPass

	s.Repo.AddUser(user)

	accessToken, err := generateJWTToken(user.UserId, accessTokenDuration)
	if err != nil {
		return "", "", err
	}

	refershToken, err := generateJWTToken(user.UserId, refershTokenDuration)
	if err != nil {
		return "", "", err
	}

	return accessToken, refershToken, nil
}

func (s *Service) LoginUser(email, password string) (string, string, error) {
	user, err := s.Repo.GetUserWhere("email", email)
	if err != nil {
		return "", "", fmt.Errorf("incorrect email or password : %v", err)
	}

	if err := comparePassword(user.Password, password); err != nil {
		return "", "", fmt.Errorf("incorrect email or password : %v", err)
	}

	accessToken, err := generateJWTToken(user.UserId, accessTokenDuration)
	if err != nil {
		return "", "", err
	}

	refershToken, err := generateJWTToken(user.UserId, refershTokenDuration)
	if err != nil {
		return "", "", err
	}

	return accessToken, refershToken, nil
}

func (s *Service) RefershToken(token string) (accessToken string, err error) {
	claims, err := validateJWTToken(token)
	if err != nil {
		return "", err
	}

	userId, err := claims.GetSubject()
	if err != nil {
		return "", err
	}

	accessToken, err = generateJWTToken(userId, accessTokenDuration)
	if err != nil {
		return "", err
	}

	return
}
