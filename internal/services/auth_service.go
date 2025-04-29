package services

import (
	"fmt"
	"time"

	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/internal/repo"
	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repo *repo.Repo
}

func NewService(repo *repo.Repo) Service {
	return Service{repo}
}

func (s Service) CreateUser(user *models.User) (string, error) {
	users := s.Repo.GetUser("email", user.Email)
	if len(users) != 0 {
		return "", fmt.Errorf("email already exists")
	}
	users = nil
	users = s.Repo.GetUser("user_name", user.UserName)
	if len(users) != 0 {
		return "", fmt.Errorf("username already exists")
	}
	user.UserId = GenerateUserID()
	hashPass, err := GenerateHashPassword(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashPass

	s.Repo.AddUser(*user)

	token, err := GenerateJWTToken(user.UserId, user.UserName)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s Service) LoginUser(email, password string) (string, error) {
	users := s.Repo.GetUser("email", email)
	if len(users) > 1 || len(users) == 0 {
		return "", fmt.Errorf("email issue")
	}
	user := users[0]
	if err := ComparePassword(user.Password, password); err != nil {
		return "", err
	}
	token, err := GenerateJWTToken(user.UserId, user.UserName)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateUserID() string {
	return uuid.New().String()
}

func GenerateHashPassword(pass string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return string(""), err
	}
	return string(hashedPass), nil
}

func ComparePassword(hashedPass, pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
	return err
}

var secrete string = "secret key"

func GenerateJWTToken(userId, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      userId,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secrete))
	return tokenString, err
}

func ValidateJWTToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secrete, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("token is invalid")
	}

	return nil
}
