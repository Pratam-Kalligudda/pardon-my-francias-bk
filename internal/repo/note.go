package repo

import (
	"errors"

	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/models"
	"gorm.io/gorm"
)

func (r *Repo) AddNote(note *models.Note) error {
	tx := r.DB.Create(&note)
	return tx.Error
}

func (r *Repo) GetAllNoteForUser(userId string) (notes []models.Note, err error) {
	tx := r.DB.Find(&notes, "user_id = ?", userId)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return notes, tx.Error
}

// func (r *Repo) UpdateNote(note models.Note) (error) {
// 	tx := r.DB.Model(&note).Update(&note)
// }
