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

// TODO: make these fields in both structs pointers and update funcs accordingly
//
//	type DocumentSearchCriteria struct {
//		Id          *uuid.UUID
//		UserID      *uuid.UUID
//		Directory   *string
//		Title       *string
//		Path        *string
//		CreatedDate *time.Time
//		Keyword     *string
//	}
type DocumentSearchCriteria struct {
	Id          string
	UserID      string
	Directory   string
	Title       string
	Path        string
	CreatedDate string
	Keyword     string
}
type UserSearchCriteria struct {
	Id            string
	UserName      string
	Editor        string
	NotesLocation string
}

//	type UpdateNoteCriteria struct {
//		UserID      *uuid.UUID
//		Directory   *string
//		Title       *string
//		Path        *string
//		CreatedDate *time.Time
//		Keyword     *string
//	}

type UpdateNoteCriteria struct {
	Id          string
	UserID      string
	Directory   string
	Title       string
	Path        string
	CreatedDate string
	Keyword     string
}

// TODO: This currently updates one note at a time. Obviously want to change this.

func (db *Database) SearchDocumentsTable(criteria DocumentSearchCriteria) ([]Doc, error) {

	var conditions []string
	var args []interface{}

	if criteria.Id != "" {
		conditions = append(conditions, "Id = ?")
		args = append(args, criteria.Id)
	}

	if criteria.UserID != "" {
		conditions = append(conditions, "UserID = ?")
		args = append(args, criteria.UserID)
	}

	if criteria.Directory != "" {
		conditions = append(conditions, "Directory = ?")
		args = append(args, criteria.Directory)
	}

	if criteria.Title != "" {
		conditions = append(conditions, "Title = ?")
		args = append(args, criteria.Title)
	}

	if criteria.Path != "" {
		conditions = append(conditions, " path = ?")
		args = append(args, criteria.Path)

	}

	if criteria.CreatedDate != "" {
		conditions = append(conditions, "CreatedDate = ?")
		args = append(args, criteria.CreatedDate)
	}

	if criteria.Keyword != "" {
		conditions = append(conditions, "keyword = ?")
		args = append(args, criteria.Keyword)
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

func (db *Database) UpdateDocumentsTable(searchResult []Doc, criteria UpdateNoteCriteria) error {

	var conditions []string
	var args []interface{}

	if criteria.Id != "" {
		conditions = append(conditions, "Id = ?")
		args = append(args, criteria.Id)
	}

	if criteria.UserID != "" {
		conditions = append(conditions, "UserID = ?")
		args = append(args, criteria.UserID)
	}

	if criteria.Directory != "" {
		conditions = append(conditions, "Directory = ?")
		args = append(args, criteria.Directory)
	}

	if criteria.Title != "" {
		conditions = append(conditions, "Title = ?")
		args = append(args, criteria.Title)
	}

	if criteria.Path != "" {
		conditions = append(conditions, " path = ?")
		args = append(args, criteria.Path)

	}

	if criteria.CreatedDate != "" {
		conditions = append(conditions, "CreatedDate = ?")
		args = append(args, criteria.CreatedDate)
	}

	if criteria.Keyword != "" {
		conditions = append(conditions, "keyword = ?")
		args = append(args, criteria.Keyword)
	}

	if len(conditions) == 0 {
		return fmt.Errorf("no search criteria provided")
	}

	for _, doc := range searchResult {

		queryArgs := make([]interface{}, len(args)+1)
		copy(queryArgs, args)
		queryArgs[len(args)] = doc.Id

		query := fmt.Sprintf("UPDATE documents SET %s WHERE id = ? ", strings.Join(conditions, " , "))
		if doc.Path != "" && criteria.Path != "" {
			if err := MoveNotes(doc.Path, criteria.Path); err != nil {
				return fmt.Errorf("Could not move notes: %w", err)
			}
		}

		_, err := db.DB.Exec(query, queryArgs...)
		if err != nil {
			return fmt.Errorf("could not execute search query: %w", err)
		}

	}

	return nil
}

// TODO: make these
func (db *Database) SearchUsersTable(criteria UserSearchCriteria) ([]User, error) {
	var conditions []string
	var args []interface{}

	if criteria.Id != "" {
		conditions = append(conditions, "id = ?")
		args = append(args, criteria.Id)
	}

	if criteria.UserName != "" {
		conditions = append(conditions, "name = ?")
		args = append(args, criteria.UserName)
	}

	if criteria.NotesLocation != "" {
		conditions = append(conditions, "notesLocation = ?")
		args = append(args, criteria.NotesLocation)
	}

	if criteria.Editor != "" {
		conditions = append(conditions, "editor = ?")
		args = append(args, criteria.Editor)
	}

	if len(conditions) == 0 {
		return nil, fmt.Errorf("no search criteria provided")
	}

	query := fmt.Sprintf("SELECT id, name, notesLocation, editor FROM users WHERE %s", strings.Join(conditions, " AND "))

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("could not execute search query: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		var idStr string
		err := rows.Scan(&idStr, &user.UserName, &user.NotesLocation, &user.Editor)
		if err != nil {
			return nil, fmt.Errorf("Could not scan row: %w", err)
		}
		user.ID, err = uuid.Parse(idStr)
		if err != nil {
			return nil, fmt.Errorf("Invalid UUID for id: %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}
func UpdateUsersTable() {}

func (db *Database) InsertUser(user *User) error {

	insert :=
		`INSERT INTO users (id, name, notesLocation, editor, configPath) 
	VALUES(?, ?, ?, ?, ?)`

	_, err := db.DB.Exec(insert, user.ID, user.UserName, user.NotesLocation, user.Editor, user.ConfigPath)
	if err != nil {
		return fmt.Errorf("Failed to insert user: %w", err)
	}

	fmt.Println("User created.")

	return nil

}

func (db *Database) InsertDoc(doc *Doc) error {
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

	result, err := db.DB.Exec(insert, // Assuming db.DB is your *sql.DB
		doc.Id.String(),
		doc.UserID.String(),
		doc.Directory,
		doc.Title,
		doc.Path,
		doc.CreatedDate.Format(time.RFC3339),
		doc.Keyword)
	if err != nil {
		return fmt.Errorf("failed to insert note: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	fmt.Printf("Document inserted successfully. Rows affected: %d\n", rowsAffected)
	return nil
}

// TODO: make this recieve uuid and delete based on that.
func (db *Database) DeleteDoc(id string) error {

	query := `DELETE FROM documents WHERE id = ?`

	result, err := db.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Failed to delete note: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Failed to get rows affected: %w", err)
	}

	fmt.Printf("Rows affected: %d\n", rowsAffected)

	if rowsAffected == 0 {
		return fmt.Errorf("No docuemnt found with id: %s", id)
	}

	return nil
}

func NewDatabase() (*Database, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not get homedir: %w", err)
	}
	db := &Database{
		Path: filepath.Join(homeDir, ".config", "doc", "doc.db"),
	}

	// Create directory if it doesn't exist
	if err := db.createDirectory(); err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}

	// Open/create the database file (only once!)
	if err := db.open(); err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	return db, nil
}

func (db *Database) createDirectory() error {
	dbDir := filepath.Dir(db.Path)
	return os.MkdirAll(dbDir, 0755)
}

func (db *Database) open() error {
	var err error
	db.DB, err = sql.Open("sqlite3", db.Path)
	if err != nil {
		return err
	}

	// Test the connection
	return db.DB.Ping()
}

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
