package service

import (
	"district/repository"
)

type AuthService struct {
	authRepository repository.AuthRepository
	userRepository repository.UserRepository
}

func NewAuthService(authRepository repository.AuthRepository) *AuthService {
	return &AuthService{authRepository: authRepository}
}

/*

func (s *AuthService) Signup(user *models.User) (*models.User, error) {
	existingUser, err := s.userRepository.GetUserByIdentification(user.Identification)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword := sha256.Sum256([]byte(user.Password))
	user.Password = string(hashedPassword[:])

	// Set created and updated timestamps
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// Save user to database
	err = s.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	// Get user by email
	user, err := s.userRepository.GetUserByIdentification(email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}

	// Check password
	err = util.CheckPassword(user.Password, password)
	if err != nil {
		return "", err
	}

	// Generate JWT token
	token, err := s.generateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) generateToken(user *models.User) (string, error) {
	// Set expiration time
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create JWT claims
	claims := &util.Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

*/
