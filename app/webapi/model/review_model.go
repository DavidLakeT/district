package model

import (
	"database/sql"
)

type Review struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	UserID    string `json:"author"`
	Content   string `json:"content"`
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
