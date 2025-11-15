package utils

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/google/uuid"
)

// setupTestDB creates a temporary database for testing
func setupTestDB(t *testing.T) *Database {
	tmpDir := t.TempDir()
	db := &Database{
		Path: filepath.Join(tmpDir, "test.db"),
	}

	if err := db.createDirectory(); err != nil {
		t.Fatalf("failed to create test directory: %v", err)
	}

	if err := db.open(); err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}

	if err := db.CreateTables(); err != nil {
		t.Fatalf("failed to create tables: %v", err)
	}

	return db
}

func TestNewDatabase(t *testing.T) {
	// if NewDatabase() does not create a db in the right place, fail.
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)

	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)

	db, err := NewDatabase()
	if err != nil {
		t.Fatalf("NewDatabase() failed: %v", err)
	}
	defer db.Close()

	expectedPath := filepath.Join(tmpDir, ".config", "doc", "doc.db")
	if db.Path != expectedPath {
		t.Errorf("expected path %s, got %s", expectedPath, db.Path)
	}

	if db.DB == nil {
		t.Error("database connection is nil")
	}
}

func TestCreateTables(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Verify users table exists
	var name string
	err := db.DB.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='users'").Scan(&name)
	if err != nil {
		t.Errorf("users table was not created: %v", err)
	}

	// Verify documents table exists
	err = db.DB.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='documents'").Scan(&name)
	if err != nil {
		t.Errorf("documents table was not created: %v", err)
	}
}

func TestInsertUser(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	user := &User{
		ID:            uuid.New(),
		UserName:      "testuser",
		NotesLocation: "/tmp/notes",
		Editor:        "vim",
		ConfigPath:    "/tmp/config",
	}

	err := db.InsertUser(user)
	if err != nil {
		t.Fatalf("InsertUser() failed: %v", err)
	}

	// Verify the user was inserted
	var userName string
	err = db.DB.QueryRow("SELECT name FROM users WHERE id = ?", user.ID.String()).Scan(&userName)
	if err != nil {
		t.Fatalf("failed to query inserted user: %v", err)
	}

	if userName != user.UserName {
		t.Errorf("expected username %s, got %s", user.UserName, userName)
	}
}

func TestInsertDoc(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// First create a user
	userID := uuid.New()
	user := &User{
		ID:            userID,
		UserName:      "testuser",
		NotesLocation: "/tmp/notes",
		Editor:        "vim",
		ConfigPath:    "/tmp/config",
	}
	db.InsertUser(user)

	// Now create a document
	doc := &Doc{
		Id:          uuid.New(),
		UserID:      userID,
		Directory:   "/tmp/notes",
		Title:       "Test Note",
		Path:        "/tmp/notes/test.md",
		CreatedDate: time.Now(),
		Keyword:     "test",
	}

	err := db.InsertDoc(doc)
	if err != nil {
		t.Fatalf("InsertDoc() failed: %v", err)
	}

	// Verify the document was inserted
	var title string
	err = db.DB.QueryRow("SELECT title FROM documents WHERE id = ?", doc.Id.String()).Scan(&title)
	if err != nil {
		t.Fatalf("failed to query inserted document: %v", err)
	}

	if title != doc.Title {
		t.Errorf("expected title %s, got %s", doc.Title, title)
	}
}

func TestDeleteDoc(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Setup: create user and document
	userID := uuid.New()
	user := &User{
		ID:            userID,
		UserName:      "testuser",
		NotesLocation: "/tmp/notes",
		Editor:        "vim",
		ConfigPath:    "/tmp/config",
	}
	db.InsertUser(user)

	docID := uuid.New()
	doc := &Doc{
		Id:          docID,
		UserID:      userID,
		Directory:   "/tmp/notes",
		Title:       "Test Note",
		Path:        "/tmp/notes/test.md",
		CreatedDate: time.Now(),
		Keyword:     "test",
	}
	db.InsertDoc(doc)

	// Test deletion
	err := db.DeleteDoc(docID.String())
	if err != nil {
		t.Fatalf("DeleteDoc() failed: %v", err)
	}

	// Verify document is deleted
	var count int
	err = db.DB.QueryRow("SELECT COUNT(*) FROM documents WHERE id = ?", docID.String()).Scan(&count)
	if err != nil {
		t.Fatalf("failed to query after deletion: %v", err)
	}

	if count != 0 {
		t.Error("document was not deleted")
	}
}

func TestDeleteDoc_NotFound(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	err := db.DeleteDoc(uuid.New().String())
	if err == nil {
		t.Error("expected error when deleting non-existent document")
	}
}

func TestSearchDocumentsTable(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Setup: create user and multiple documents
	userID := uuid.New()
	user := &User{
		ID:            userID,
		UserName:      "testuser",
		NotesLocation: "/tmp/notes",
		Editor:        "vim",
		ConfigPath:    "/tmp/config",
	}
	db.InsertUser(user)

	docs := []*Doc{
		{
			Id:          uuid.New(),
			UserID:      userID,
			Title:       "First Note",
			Path:        "/tmp/notes/first.md",
			CreatedDate: time.Now(),
			Keyword:     "golang",
		},
		{
			Id:          uuid.New(),
			UserID:      userID,
			Title:       "Second Note",
			Path:        "/tmp/notes/second.md",
			CreatedDate: time.Now(),
			Keyword:     "testing",
		},
	}

	for _, doc := range docs {
		db.InsertDoc(doc)
	}

	tests := []struct {
		name     string
		criteria DocumentSearchCriteria
		wantLen  int
	}{
		{
			name:     "search by keyword",
			criteria: DocumentSearchCriteria{Keyword: "golang"},
			wantLen:  1,
		},
		{
			name:     "search by title",
			criteria: DocumentSearchCriteria{Title: "First Note"},
			wantLen:  1,
		},
		{
			name:     "search by user ID",
			criteria: DocumentSearchCriteria{UserID: userID.String()},
			wantLen:  2,
		},
		{
			name:     "search by ID",
			criteria: DocumentSearchCriteria{Id: docs[0].Id.String()},
			wantLen:  1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results, err := db.SearchDocumentsTable(test.criteria)
			if err != nil {
				t.Fatalf("SearchDocumentsTable() failed: %v", err)
			}

			if len(results) != test.wantLen {
				t.Errorf("expected %d results, got %d", test.wantLen, len(results))
			}
		})
	}
}

func TestSearchDocumentsTable_NoCriteria(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	_, err := db.SearchDocumentsTable(DocumentSearchCriteria{})
	if err == nil {
		t.Error("expected error when no search criteria provided")
	}
}

func TestSearchUsersTable(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Setup: create multiple users
	users := []*User{
		{
			ID:            uuid.New(),
			UserName:      "alice",
			NotesLocation: "/home/alice/notes",
			Editor:        "vim",
			ConfigPath:    "/home/alice/.config",
		},
		{
			ID:            uuid.New(),
			UserName:      "bob",
			NotesLocation: "/home/bob/notes",
			Editor:        "emacs",
			ConfigPath:    "/home/bob/.config",
		},
	}

	for _, user := range users {
		db.InsertUser(user)
	}

	tests := []struct {
		name     string
		criteria UserSearchCriteria
		wantLen  int
	}{
		{
			name:     "search by username",
			criteria: UserSearchCriteria{UserName: "alice"},
			wantLen:  1,
		},
		{
			name:     "search by editor",
			criteria: UserSearchCriteria{Editor: "vim"},
			wantLen:  1,
		},
		{
			name:     "search by ID",
			criteria: UserSearchCriteria{Id: users[0].ID.String()},
			wantLen:  1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results, err := db.SearchUsersTable(test.criteria)
			if err != nil {
				t.Fatalf("SearchUsersTable() failed: %v", err)
			}

			if len(results) != test.wantLen {
				t.Errorf("expected %d results, got %d", test.wantLen, len(results))
			}
		})
	}
}

func TestSearchUsersTable_NoCriteria(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	_, err := db.SearchUsersTable(UserSearchCriteria{})
	if err == nil {
		t.Error("expected error when no search criteria provided")
	}
}

func TestUpdateDocumentsTable(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Setup: create user and document
	userID := uuid.New()
	user := &User{
		ID:            userID,
		UserName:      "testuser",
		NotesLocation: "/tmp/notes",
		Editor:        "vim",
		ConfigPath:    "/tmp/config",
	}
	db.InsertUser(user)

	doc := &Doc{
		Id:          uuid.New(),
		UserID:      userID,
		Title:       "Original Title",
		Path:        "/tmp/notes/original.md",
		CreatedDate: time.Now(),
		Keyword:     "original",
	}
	db.InsertDoc(doc)

	// Search for the document
	searchResults, err := db.SearchDocumentsTable(DocumentSearchCriteria{
		Id: doc.Id.String(),
	})
	if err != nil {
		t.Fatalf("SearchDocumentsTable() failed: %v", err)
	}

	// Update the document
	updateCriteria := UpdateNoteCriteria{
		Title:   "Updated Title",
		Keyword: "updated",
	}

	err = db.UpdateDocumentsTable(searchResults, updateCriteria)
	if err != nil {
		t.Fatalf("UpdateDocumentsTable() failed: %v", err)
	}

	// Verify the update
	var title, keyword string
	err = db.DB.QueryRow("SELECT title, keyword FROM documents WHERE id = ?", doc.Id.String()).Scan(&title, &keyword)
	if err != nil {
		t.Fatalf("failed to query updated document: %v", err)
	}

	if title != "Updated Title" {
		t.Errorf("expected title 'Updated Title', got %s", title)
	}

	if keyword != "updated" {
		t.Errorf("expected keyword 'updated', got %s", keyword)
	}
}

func TestClose(t *testing.T) {
	db := setupTestDB(t)

	err := db.Close()
	if err != nil {
		t.Errorf("Close() failed: %v", err)
	}

	// Verify database is closed by trying to query
	err = db.DB.Ping()
	if err == nil {
		t.Error("expected error after closing database")
	}
}

func TestClose_NilDB(t *testing.T) {
	db := &Database{}
	err := db.Close()
	if err != nil {
		t.Errorf("Close() with nil DB failed: %v", err)
	}
}

func TestDatabase_Exists(t *testing.T) {
	// Create a temporary directory for our test
	tempDir := t.TempDir()

	// Create a test database path
	dbPath := filepath.Join(tempDir, "test.db")

	// Create the database struct
	db := &Database{
		Path: dbPath,
	}

	// Test 1: Database file doesn't exist yet
	if db.Exists() {
		t.Error("Expected Exists() to return false, but got true")
	}

	// Create an empty file at that path
	file, err := os.Create(dbPath)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	file.Close()

	// Test 2: Now the database file exists
	if !db.Exists() {
		t.Error("Expected Exists() to return true, but got false")
	}
}
