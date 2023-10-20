package service

import (
	controller "district/controller/request"
	"district/model"
	dto "district/model/dto"
	"district/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.userRepository.CreateUser(user)
}

func (s *UserService) GetUserByIdentification(identification int) (*dto.UserDTO, error) {
	user, err := s.userRepository.GetUserByIdentification(identification)
	if err != nil {
		return nil, err
	}

	return dto.ConvertToUserDTO(user), nil
}

func (s *UserService) UpdateUser(identification int, request *controller.UpdateUserRequest) error {
	user, err := s.userRepository.GetUserByIdentification(identification)
	if err != nil {
		return err
	}

	switch {
	case request.Email != "":
		user.Email = request.Email
	case request.Username != "":
		user.Username = request.Username
	case request.Password != "":
		user.Password = request.Password
	case request.Address != "":
		user.Address = request.Address
	case request.IsAdmin != nil:
		user.IsAdmin = request.IsAdmin
	}

	return s.userRepository.UpdateUser(user)
}

func (s *UserService) DeleteUserByIdentification(identification int) error {
	return s.userRepository.DeleteUser(identification)
}
