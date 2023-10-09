package service

import (
	"district/models"
	"district/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepo}
}

func (s *UserService) GetUserByIdentification(identification int) (*models.User, error) {
	return s.userRepository.GetUserByIdentification(identification)
}
