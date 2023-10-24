package service

import (
	controller "district/controller/request"
	"district/model"
	dto "district/model/dto"
	repository "district/repository/handler"
)

type UserService struct {
	repositoryPool *repository.RepositoryPool
}

func NewUserService(repositoryPool *repository.RepositoryPool) *UserService {
	return &UserService{repositoryPool: repositoryPool}
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.repositoryPool.UserRepository.CreateUser(user)
}

func (s *UserService) GetUserByIdentification(identification int) (*dto.UserDTO, error) {
	user, err := s.repositoryPool.UserRepository.GetUserByIdentification(identification)
	if err != nil {
		return nil, err
	}

	return dto.ConvertToUserDTO(user), nil
}

func (s *UserService) UpdateUser(identification int, request *controller.UpdateUserRequest) error {
	user, err := s.repositoryPool.UserRepository.GetUserByIdentification(identification)
	if err != nil {
		return err
	}

	switch {
	case request.Email != nil:
		user.Email = *request.Email
		fallthrough
	case request.Username != nil:
		user.Username = *request.Username
		fallthrough
	case request.Password != nil:
		user.Password = *request.Password
		fallthrough
	case request.Address != nil:
		user.Address = *request.Address
		fallthrough
	case request.Balance != nil:
		user.Balance = *request.Balance
		fallthrough
	case request.IsAdmin != nil:
		user.IsAdmin = *request.IsAdmin
	}

	return s.repositoryPool.UserRepository.UpdateUser(user)
}

func (s *UserService) DeleteUserByIdentification(identification int) error {
	return s.repositoryPool.UserRepository.DeleteUser(identification)
}
