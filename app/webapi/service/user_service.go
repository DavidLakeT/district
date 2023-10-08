package service

import (
	"errors"

	"district/models"
	"district/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepo}
}

func (s *UserService) CreateUser(user *models.User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	return s.userRepository.CreateUser(user)
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.userRepository.GetUserByID(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	return s.userRepository.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
	if id <= 0 {
		return errors.New("invalid user ID")
	}
	return s.userRepository.DeleteUser(id)
}
