package services

import "github.com/Pratam-Kalligudda/pardon-my-francias-bk/internal/repo"

type Service struct {
	Repo *repo.Repo
}

func NewService(repo *repo.Repo) Service {
	return Service{repo}
}
