package repository

import (
	"be-skillacademy-final/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) Add(user model.User) error {
	if result := u.db.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (u *userRepository) CheckAvail(user model.User) error {
	result := u.db.Model(model.User{}).Where("email = ? AND password = ?", user.Email, user.Password).First(&user);
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}
