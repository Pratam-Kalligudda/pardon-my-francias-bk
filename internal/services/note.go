package services

import (
	"github.com/Pratam-Kalligudda/pardon-my-francias-bk/models"
)

func (s *Service) CreateNote(note *models.Note, token string) error {
	claims, err := validateJWTToken(token)
	if err != nil {
		return err
	}
	userID, err := claims.GetSubject()
	if err != nil {
		return err
	}

	note.UserID = userID
	err = s.Repo.AddNote(note)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetNotesOfUser(token string) ([]models.Note, error) {
	claims, err := validateJWTToken(token)
	if err != nil {
		return nil, err
	}
	userID, err := claims.GetSubject()
	if err != nil {
		return nil, err
	}

	notes, err := s.Repo.GetAllNoteForUser(userID)
	if err != nil {
		return nil, err
	}

	return notes, nil
}
