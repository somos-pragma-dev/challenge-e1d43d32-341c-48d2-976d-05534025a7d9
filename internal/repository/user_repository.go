package repository

import (
	"errors"
	"gorm.io/gorm"
	"internal/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUsers() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err!= nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err!= nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) DeleteUser(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}