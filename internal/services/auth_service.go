package services

import (
	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/internal/repo"
	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repo *repo.Repo
}

func NewService(repo *repo.Repo) Service {
	return Service{repo}
}

func (s Service) CreateUser(user *models.User) error {
	user.UserId = GenerateUserID()
	hashPass, err := GenerateHashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPass

	s.Repo.AddUser(*user)

	return nil
}

func (s Service) LoginUser(email, password string) error {
	hashPass := s.Repo.GetPassword(email)
	if err := ComparePassword(hashPass, password); err != nil {
		return err
	}
	return nil
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
