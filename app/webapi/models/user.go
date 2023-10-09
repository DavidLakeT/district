package models

import "time"

type User struct {
	Identification int     `json:"identification"`
	Email          string  `json:"email"`
	Username       string  `json:"username"`
	Password       string  `json:"password"`
	Address        string  `json:"address"`
	Balance        float64 `json:"balance"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeleteAt       time.Time
}
