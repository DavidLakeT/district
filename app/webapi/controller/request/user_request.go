package controller

type CreateUserRequest struct {
	Identification int    `json:"identification"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Address        string `json:"address"`
	IsAdmin        *bool  `json:"isAdmin"`
}

type UpdateUserRequest struct {
	Email    *string  `json:"email"`
	Username *string  `json:"username"`
	Password *string  `json:"password"`
	Address  *string  `json:"address"`
	Balance  *float64 `json:"balance"`
	IsAdmin  *bool    `json:"isAdmin"`
}
