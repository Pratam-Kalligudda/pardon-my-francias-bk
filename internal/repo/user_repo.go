package repo

import (
	"fmt"

	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/models"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return Repo{db}
}

func (r *Repo) AddUser(user models.User) {
	r.DB.Create(&user)
	var users []models.User
	r.DB.Find(&users)
	fmt.Println("Get Users:", users)
}

func (r *Repo) GetPassword(email string) string {
	var users []models.User
	r.DB.Find(&users, "email", email)
	if len(users) == 0 {
		return ""
	}
	return users[0].Password
}
