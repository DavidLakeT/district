package auth

// LoginResponse represents the response body for the login endpoint
type LoginResponse struct {
	Token string `json:"token"`
}

// RegisterResponse represents the response body for the register endpoint
type RegisterResponse struct {
	Message string `json:"message"`
}
