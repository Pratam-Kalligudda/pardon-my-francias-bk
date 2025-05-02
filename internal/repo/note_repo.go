package repo

import (
	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/models"
)

func (r *Repo) AddNote(note models.Note) (*models.Note, error) {
	tx := r.DB.Create(&note)
	return &note, tx.Error
}

// func (r *Repo) UpdateNote(note models.Note) (error) {
// 	tx := r.DB.Model(&note).Update(&note)
// }
