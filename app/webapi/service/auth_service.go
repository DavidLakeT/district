package service

import (
	repository "district/repository/handler"
	"encoding/base64"
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
		return "", fmt.Errorf("there is no account with the email address you provided.")
	}

	if !CheckPasswordHash(password, user.Password) {
		return "", fmt.Errorf("your password doesn't match. Please try again.")
	}

	token := fmt.Sprintf("%d:%s:%s:%v", user.Identification, user.Username, user.Email, user.IsAdmin)
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
