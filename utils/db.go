package utils

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Path string
	DB   *sql.DB
}

type SearchCriteria struct {
	Keyword     string
	Title       string
	Location    string
	CreatedDate string
}

func (db *Database) Search(criteria SearchCriteria) ([]Note, error) {
	var conditions []string
	var args []interface{}

	if criteria.Keyword != "" {
		conditions = append(conditions, "keyword = ?")
		args = append(args, criteria.Keyword)
	}

	if criteria.Title != "" {
		conditions = append(conditions, "title = ?")
		args = append(args, criteria.Title)
	}

	if criteria.Location != "" {
		conditions = append(conditions, "location = ?")
		args = append(args, criteria.Location)

	}

	if criteria.CreatedDate != "" {
		conditions = append(conditions, "created_date = ?")
		args = append(args, criteria.CreatedDate)

	}

	if len(conditions) == 0 {
		return nil, fmt.Errorf("no search criteria provided")
	}

	query := fmt.Sprintf("SELECT id, keyword, title, location, created_date FROM documents WHERE %s", strings.Join(conditions, " AND "))

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("could not execute search query: %w", err)
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var note Note
		var CreatedDateStr string
		var idStr string
		err := rows.Scan(&idStr, &note.Keyword, &note.Title, &note.Location, &CreatedDateStr)
		if err != nil {
			return nil, fmt.Errorf("Could not scan row: %w", err)
		}
		note.Id, err = uuid.Parse(idStr)
		if err != nil {
			return nil, fmt.Errorf("Invalid UUID for id: %w", err)
		}
		if CreatedDateStr != "" {
			note.CreatedDate, err = time.Parse(time.RFC3339, CreatedDateStr)
		}
		notes = append(notes, note)
	}
	return notes, nil
}

func InsertNote(db *sql.DB, note *Note) error {
	fmt.Printf("DEBUG: InsertNote started\n")
	insert := `
	INSERT INTO documents (id, user_id, title, extension, location, created_date, keyword) 
	VALUES(?, ?, ?, ?, ?, ?, ?)
	`
	rows, err := db.Exec(insert,
		note.Id.String(),
		note.User_id.String(),
		note.Title,
		note.Extension,
		note.Location,
		note.CreatedDate.Format(time.RFC3339),
		note.Keyword)

	if err != nil {
		return fmt.Errorf("Failed to insert note: %w", err)
	}

	fmt.Printf("DEBUG: Exec succeeded. Rows affected: %d\n", rows)

	return nil
}

// NewDatabase creates a new database at the specified path
func NewDatabase() (*Database, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not get homedir: %w", err)
	}
	db := &Database{
		Path: filepath.Join(homeDir, ".config", "doc", "doc.db"),
	}
	db.open()
	// fmt.Println("DB Path = ", db.Path)
	// fmt.Println("DB Driver = ", db.DB)

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
