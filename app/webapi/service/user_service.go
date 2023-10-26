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

	if request.Email != nil {
		user.Email = *request.Email
	}
	if request.Username != nil {
		user.Username = *request.Username
	}
	if request.Password != nil {
		hashedPassword, err := HashPassword(*request.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}
	if request.Address != nil {
		user.Address = *request.Address
	}
	if request.Balance != nil {
		user.Balance = *request.Balance
	}
	if request.IsAdmin != nil {
		user.IsAdmin = *request.IsAdmin
	}

	return s.repositoryPool.UserRepository.UpdateUser(user)
}

func (s *UserService) DeleteUserByIdentification(identification int) error {
	return s.repositoryPool.UserRepository.DeleteUser(identification)
}
