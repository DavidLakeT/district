package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbPath = "./users.db"
)

// User represents a user in the database
type User struct {
	ID       int
	Username string
	Password string
}

// Connect creates a connection to the Users database
func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil
}

// CreateUserTable creates the Users table in the database
func CreateUserTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Users (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		Username TEXT NOT NULL,
		Password TEXT NOT NULL
	)`)
	if err != nil {
		return fmt.Errorf("failed to create Users table: %v", err)
	}

	return nil
}

// AddUsers adds some sample users to the database
func AddUsers(db *sql.DB) error {
	users := []User{
		{Username: "user1", Password: "password1"},
		{Username: "user2", Password: "password2"},
		{Username: "user3", Password: "password3"},
	}

	for _, user := range users {
		_, err := db.Exec("INSERT INTO Users (Username, Password) VALUES (?, ?)", user.Username, user.Password)
		if err != nil {
			return fmt.Errorf("failed to add user %s: %v", user.Username, err)
		}
	}

	return nil
}
