package utils

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Path string
	DB   *sql.DB
}

// TODO: finish these
// func (db *Database) UpdateNote(noteID int, title, content string) error
// func (db *Database) GetNoteByPath(filePath string) (*Note, error)
// func (db *Database) InsertNote(userID int, title, content string) (int64, error)

// NewDatabase creates a new database at the specified path
func NewDatabase(dbPath string) (*Database, error) {
	db := &Database{
		Path: dbPath,
	}

	// Create directory if it doesn't exist
	if err := db.createDirectory(); err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}

	// Open/create the database file
	if err := db.open(); err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	return db, nil
}

// createDirectory creates the directory for the database if it doesn't exist
func (db *Database) createDirectory() error {
	dbDir := filepath.Dir(db.Path)
	return os.MkdirAll(dbDir, 0755)
}

// open opens the database connection
func (db *Database) open() error {
	var err error
	db.DB, err = sql.Open("sqlite3", db.Path)
	if err != nil {
		return err
	}

	// Test the connection
	return db.DB.Ping()
}

// CreateTables creates the users and notes tables
func (db *Database) CreateTables() error {
	if err := db.createUsersTable(); err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	if err := db.createNotesTable(); err != nil {
		return fmt.Errorf("failed to create notes table: %w", err)
	}

	return nil
}

// createUsersTable creates the users table
func (db *Database) createUsersTable() error {
	query := `
	 CREATE TABLE IF NOT EXISTS users (
		 id TEXT PRIMARY KEY, -- UUID as TEXT
		 name TEXT NOT NULL,
		 notes_location TEXT NOT NULL
	 );
	`
	_, err := db.DB.Exec(query)
	return err
}

// createNotesTable creates the notes table
func (db *Database) createNotesTable() error {
	query := `
	 CREATE TABLE IF NOT EXISTS documents (
		 id TEXT PRIMARY KEY,
		 user_id TEXT NOT NULL,
		 title TEXT,
		 extension TEXT,
		 location TEXT,
		 full_name TEXT,
		 created_date TEXT, -- ISO 8601 format (used for date/time)
		 keyword TEXT,
		 inode INTEGER,
		 FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	 );

	`
	_, err := db.DB.Exec(query)
	return err
}

// Exists checks if the database file exists
func (db *Database) Exists() bool {
	_, err := os.Stat(db.Path)
	return !os.IsNotExist(err)
}

// Close closes the database connection
func (db *Database) Close() error {
	if db.DB != nil {
		return db.DB.Close()
	}
	return nil
}
