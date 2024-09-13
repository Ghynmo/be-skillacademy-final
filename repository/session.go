package repository

import (
	"be-skillacademy-final/model"

	"gorm.io/gorm"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)
}

type sessionsRepoImpl struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (s *sessionsRepoImpl) AddSessions(session model.Session) error {
	if result := s.db.Create(&session); result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	if result := s.db.Where("token = ?", token).Delete(&model.Session{}); result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	result := s.db.Where("email = ?", session.Email).Updates(&session); 
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) SessionAvailName(email string) error {
	result := s.db.Model(model.Session{}).Where("email = ?", email).First(&model.Session{}); 
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	session := model.Session{}
	result := s.db.Model(model.Session{}).Where("token = ?", token).First(&model.Session{}).Scan(&session); 
	if result.Error != nil {
		return model.Session{}, result.Error
	}
	return session, nil // TODO: replace this
}
