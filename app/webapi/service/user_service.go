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

func (s *UserService) CreateUser(user *model.User) error {
	return s.userRepository.CreateUser(user)
}

func (s *UserService) GetUserByIdentification(identification int) (*model.User, error) {
	return s.userRepository.GetUserByIdentification(identification)
}

func (s *UserService) UpdateUser(user *model.User) error {
	return s.userRepository.UpdateUser(user)
}

func (s *UserService) DeleteUserByIdentification(identification int) error {
	return s.userRepository.DeleteUser(identification)
}
