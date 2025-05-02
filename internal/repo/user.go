package repo

import (
	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/models"
)

func (r *Repo) AddUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *Repo) CheckIfUserExists(keyword string, value string) error {
	var user models.User
	return r.DB.First(&user, keyword+" = ?", value).Error
}

func (r *Repo) GetUserWhere(keyword string, value string) (*models.User, error) {
	var user models.User
	tx := r.DB.First(&user, keyword+" = ?", value)
	return &user, tx.Error
}

func (r *Repo) GetAllUser() ([]models.User, error) {
	var users []models.User
	tx := r.DB.Find(&users)
	return users, tx.Error
}
