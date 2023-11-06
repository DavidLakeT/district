package service

import (
	repository "district/repository/handler"
	"encoding/base64"
	"fmt"
)

type AuthService struct {
	repositoryPool *repository.RepositoryPool
}

func NewAuthService(repositoryPool *repository.RepositoryPool) *AuthService {
	return &AuthService{repositoryPool: repositoryPool}
}

func (as *AuthService) Login(email, password string) (string, error) {
	user, err := as.repositoryPool.AuthRepository.VerifyLogin(email, password)
	if err != nil {
		fmt.Println("error:", err)
		return "", fmt.Errorf("there is no account with the credentials you provided.")
	}

	token := fmt.Sprintf("%d:%s:%s:%v", user.Identification, user.Username, user.Email, user.IsAdmin)
	encodedToken := base64.StdEncoding.EncodeToString([]byte(token))

	return encodedToken, nil
}
