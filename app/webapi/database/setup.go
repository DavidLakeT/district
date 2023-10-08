package database

import (
	"database/sql"
)

func SetupDatabase(db *sql.DB) error {
	_, err := db.Exec(`DROP TABLE IF EXISTS users;`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			username TEXT NOT NULL,
			password TEXT NOT NULL
		);
	`)
	if err != nil {
		return err
	}

	return nil
}
