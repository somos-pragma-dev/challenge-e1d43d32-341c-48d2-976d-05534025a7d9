package service

import (
	"errors"
	"internal/model"
	"internal/repository"
	"pkg/validation"
)

type UserService struct {
	userRepo repository.UserRepository
	validator *validation.Validator
}

func NewUserService(userRepo repository.UserRepository, validator *validation.Validator) *UserService {
	return &UserService{userRepo: userRepo, validator: validator}
}

func (s *UserService) CreateUser(user *model.User) error {
	if err := s.validator.Validate(user); err!= nil {
		return err
	}
	return s.userRepo.CreateUser(user)
}

func (s *UserService) GetUsers() ([]model.User, error) {
	return s.userRepo.GetUsers()
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *UserService) UpdateUser(user *model.User) error {
	if err := s.validator.Validate(user); err!= nil {
		return err
	}
	return s.userRepo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.DeleteUser(id)
}