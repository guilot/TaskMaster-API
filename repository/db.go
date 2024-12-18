package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite Driver
)

// InitDB function
func InitDB(dbName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}

	// Tables creation
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`
	createTaskTable := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		status TEXT NOT NULL DEFAULT 'pendiente',
		due_date TEXT,
		user_id INTEGER NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`

	// ignore result of creation with _, because it's not important
	// Just check if an error occurs with err as the second return value
	_, err = db.Exec(createUserTable)
	if err != nil {
		log.Fatalf("Error creating table users: %v", err)
		return nil, err
	}

	_, err = db.Exec(createTaskTable)
	if err != nil {
		log.Fatalf("Error creating table tasks: %v", err)
		return nil, err
	}

	return db, nil
}
