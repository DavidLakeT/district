package auth

// LoginRequest represents the request body for the login endpoint
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterRequest represents the request body for the register endpoint
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
