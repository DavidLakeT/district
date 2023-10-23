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
	return s.repositoryPool.GetUserRepository().CreateUser(user)
}

func (s *UserService) GetUserByIdentification(identification int) (*dto.UserDTO, error) {
	user, err := s.repositoryPool.GetUserRepository().GetUserByIdentification(identification)
	if err != nil {
		return nil, err
	}

	return dto.ConvertToUserDTO(user), nil
}

func (s *UserService) UpdateUser(identification int, request *controller.UpdateUserRequest) error {
	user, err := s.repositoryPool.GetUserRepository().GetUserByIdentification(identification)
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

	return s.repositoryPool.GetUserRepository().UpdateUser(user)
}

func (s *UserService) DeleteUserByIdentification(identification int) error {
	return s.repositoryPool.GetUserRepository().DeleteUser(identification)
}
