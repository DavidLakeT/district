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

func (us *UserService) GetAllUsers() ([]*dto.UserDTO, error) {
	users, err := us.repositoryPool.UserRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	userDTOs := make([]*dto.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = dto.ConvertToUserDTO(user)
	}

	return userDTOs, nil
}

func (us *UserService) CreateUser(request *controller.CreateUserRequest) (*dto.UserDTO, error) {
	hashedPassword, err := HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	admin := false
	if request.IsAdmin != nil {
		admin = *request.IsAdmin
	}

	user := model.User{
		Identification: request.Identification,
		Email:          request.Email,
		Username:       request.Username,
		Password:       hashedPassword,
		Address:        request.Address,
		IsAdmin:        admin,
	}

	err = us.repositoryPool.UserRepository.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return dto.ConvertToUserDTO(&user), nil
}

func (us *UserService) GetUserByIdentification(identification int) (*dto.UserDTO, error) {
	user, err := us.repositoryPool.UserRepository.GetUserByIdentification(identification)
	if err != nil {
		return nil, err
	}

	return dto.ConvertToUserDTO(user), nil
}

func (us *UserService) UpdateUser(identification int, request *controller.UpdateUserRequest) error {
	user, err := us.repositoryPool.UserRepository.GetUserByIdentification(identification)
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

	return us.repositoryPool.UserRepository.UpdateUser(user)
}

func (us *UserService) DeleteUserByIdentification(identification int) error {
	return us.repositoryPool.UserRepository.DeleteUser(identification)
}

func (us *UserService) DeleteUsersTable() error {
	if err := us.repositoryPool.UserRepository.DeleteUsersTable(); err != nil {
		return err
	}

	return nil
}
