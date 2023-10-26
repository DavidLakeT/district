package model

import "district/model"

type UserDTO struct {
	Identification int     `json:"identification"`
	Email          string  `json:"email"`
	Username       string  `json:"username"`
	Address        string  `json:"address"`
	Balance        float64 `json:"balance"`
	IsAdmin        bool    `json:"is_admin,omitempty"`
}

func ConvertToUserDTO(user *model.User) *UserDTO {
	return &UserDTO{
		Identification: user.Identification,
		Email:          user.Email,
		Username:       user.Username,
		Address:        user.Address,
		Balance:        user.Balance,
		IsAdmin:        user.IsAdmin,
	}
}
