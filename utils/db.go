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
	Path        string
	CreatedDate string
}

func (db *Database) Search(criteria SearchCriteria) ([]Doc, error) {
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

	if criteria.Path != "" {
		conditions = append(conditions, " path= ?")
		args = append(args, criteria.Path)

	}

	if criteria.CreatedDate != "" {
		conditions = append(conditions, "created_date = ?")
		args = append(args, criteria.CreatedDate)

	}

	if len(conditions) == 0 {
		return nil, fmt.Errorf("no search criteria provided")
	}

	query := fmt.Sprintf("SELECT id, keyword, title, path, created_date FROM documents WHERE %s", strings.Join(conditions, " AND "))

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("could not execute search query: %w", err)
	}
	defer rows.Close()

	var notes []Doc
	for rows.Next() {
		var note Doc
		var CreatedDateStr string
		var idStr string
		err := rows.Scan(&idStr, &note.Keyword, &note.Title, &note.Path, &CreatedDateStr)
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

func InsertUser(db *sql.DB, user *User) error {

	insert :=
		`INSERT INTO users (id, name, notesLocation, editor, configPath) 
	VALUES(?, ?, ?, ?, ?)`

	_, err := db.Exec(insert, user.ID, user.UserName, user.NotesLocation, user.Editor, user.ConfigPath)
	if err != nil {
		return fmt.Errorf("Failed to insert user: %w", err)
	}

	fmt.Println("User created.")

	return nil

}

func InsertNote(db *sql.DB, doc *Doc) error {

	fmt.Println("ID: ", doc.Id)
	fmt.Println("UserID", doc.UserID)
	fmt.Println("Title: ", doc.Title)
	fmt.Println("Path: ", doc.Path)
	fmt.Println("Create Date: ", doc.CreatedDate)
	fmt.Println("Keyword: ", doc.Keyword)

	insert := `
	INSERT INTO documents (id, userID, directory, title, path, created_date, keyword) 
	VALUES(?, ?, ?, ?, ?, ?, ?)
	`
	rows, err := db.Exec(insert,
		doc.Id.String(),
		doc.UserID.String(),
		doc.Directory,
		doc.Title,
		doc.Path,
		doc.CreatedDate.Format(time.RFC3339),
		doc.Keyword)

	if err != nil {
		return fmt.Errorf("Failed to insert note: %w", err)
	}

	fmt.Printf("Exec succeeded. Rows affected: %d\n", rows)

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

	if err := db.createDocumentsTable(); err != nil {
		return fmt.Errorf("failed to create notes table: %w", err)
	}

	return nil
}

func (db *Database) createUsersTable() error {
	query := `
	 CREATE TABLE IF NOT EXISTS users (
		 id TEXT PRIMARY KEY, -- UUID as TEXT
		 name TEXT NOT NULL,
		 notesLocation TEXT NOT NULL,
		 editor TExT NOT NULL,
		 configPath TEXT NOT NULL
	 );
	`
	_, err := db.DB.Exec(query)
	return err
}

func (db *Database) createDocumentsTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS documents (
		id TEXT PRIMARY KEY,
		userID TEXT NOT NULL,
		directory TEXT,
		title TEXT,
		path TEXT,
		created_date TEXT, -- ISO 8601 format (used for date/time)
		keyword TEXT,
		FOREIGN KEY (userID) REFERENCES users(id) ON DELETE CASCADE
	 );
	`
	_, err := db.DB.Exec(query)
	return err
}

func (db *Database) Exists() bool {
	_, err := os.Stat(db.Path)
	return !os.IsNotExist(err)
}

func (db *Database) Close() error {
	if db.DB != nil {
		return db.DB.Close()
	}
	return nil
}
