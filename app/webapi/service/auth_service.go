package service

import (
	repository "district/repository/handler"
	"encoding/base64"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repositoryPool *repository.RepositoryPool
}

func NewAuthService(repositoryPool *repository.RepositoryPool) *AuthService {
	return &AuthService{repositoryPool: repositoryPool}
}

func (as *AuthService) Login(email, password string) (string, error) {
	user, err := as.repositoryPool.UserRepository.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if !CheckPasswordHash(password, user.Password) {
		return "", errors.New("wrong credentials")
	}

	admin := true
	if user.IsAdmin == nil {
		admin = false
	}

	token := fmt.Sprintf("%d:%s:%s:%v", user.Identification, user.Username, user.Email, admin)
	encodedToken := base64.StdEncoding.EncodeToString([]byte(token))

	return encodedToken, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
