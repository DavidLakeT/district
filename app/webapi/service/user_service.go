package service

import (
	"district/model"
	"district/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepo}
}

func (s *UserService) GetUserByIdentification(identification int) (*model.User, error) {
	return s.userRepository.GetUserByIdentification(identification)
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.userRepository.CreateUser(user)
}
