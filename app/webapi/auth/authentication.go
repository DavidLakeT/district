package auth

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"district/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// JWTClaims represents the claims in the JWT token
type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// AuthHandler represents the authentication handler
type AuthHandler struct {
	DB       *sql.DB
	Secret   []byte
	TokenTTL time.Duration
}

// Login handles the login endpoint
func (h *AuthHandler) Login(c echo.Context) error {
	// Parse the request body
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	// Get the user from the database
	user := new(models.User)
	row := h.DB.QueryRow("SELECT id, username, password FROM users WHERE username = ?", req.Username)
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid username or password")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "database error")
	}

	// Check if the password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid username or password")
	}

	// Create the JWT token
	claims := &JWTClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(h.TokenTTL).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(h.Secret)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create token")
	}

	// Return the token
	res := &LoginResponse{
		Token: tokenString,
	}
	return c.JSON(http.StatusOK, res)
}

// Register handles the register endpoint
func (h *AuthHandler) Register(c echo.Context) error {
	// Parse the request body
	req := new(RegisterRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to hash password")
	}

	// Insert the user into the database
	result, err := h.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", req.Username, hashedPassword)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "database error")
	}

	// Check if the user was inserted successfully
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "database error")
	}
	if rowsAffected == 0 {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to insert user")
	}

	// Return a success message
	res := &RegisterResponse{
		Message: "user registered successfully",
	}
	return c.JSON(http.StatusOK, res)
}

// Middleware function to check if the JWT token is valid
func (h *AuthHandler) JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the JWT token from the Authorization header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing Authorization header")
		}
		tokenString := authHeader[len("Bearer "):]

		// Parse the JWT token
		claims := &JWTClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return h.Secret, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token signature")
			}
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}
		if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		// Set the username in the context
		c.Set("username", claims.Username)

		// Call the next handler
		return next(c)
	}
}

// GetUser returns the user from the database with the given ID
func (h *AuthHandler) GetUser(id int) (*models.User, error) {
	user := new(models.User)
	row := h.DB.QueryRow("SELECT id, username, password FROM users WHERE id = ?", id)
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("database error")
	}
	return user, nil
}
