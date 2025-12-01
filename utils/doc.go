package utils

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type Doc struct {
	Id          uuid.UUID
	UserID      uuid.UUID
	Directory   string
	Title       string
	Path        string
	CreatedDate time.Time
	Keyword     string
}

func NewDoc(title, keyword string) (*Doc, error) {

	doc := &Doc{}

	doc.setID()
	doc.setUserID()
	if err := doc.setTitle(title); err != nil {
		return nil, err
	}
	if err := doc.setDirectory(title); err != nil {
		return nil, err
	}
	doc.setPath(doc.Directory, title)
	doc.setCreatedDate()
	doc.setKeyword(keyword)

	return doc, nil

}

func CreateDocFile(title, keyword string) error {

	doc, err := NewDoc(title, keyword)
	if err != nil {
		return fmt.Errorf("Could not create note: %w", err)
	}

	frontmatter := SetFrontmatter(doc)

	err = os.WriteFile(doc.Path, []byte(frontmatter), 0644)
	if err != nil {
		return fmt.Errorf("Could not write to new doc file: %w", err)
	}

	return nil
}

// setPath() always appends the provided directory and title to the notes location the user defined upon setup.
// ex: setPath(math, slope.md) yields path/to/notes/location/math/slope.md
func (doc *Doc) setPath(directory, title string) {
	user, err := GetUserConfig()
	if err != nil {
		fmt.Printf("Could not get user config: %s", err)
	}
	notesLocation := user.NotesLocation
	doc.Path = filepath.Join(notesLocation, directory, title)
}

func (doc *Doc) setID() {
	doc.Id = uuid.New()
}

func (doc *Doc) GetID() uuid.UUID {
	return doc.Id
}

func (doc *Doc) setUserID() error {
	homeDir, _ := os.UserHomeDir()
	dbPath := filepath.Join(homeDir, ".config/doc/doc.db")

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	var userID string
	query := `SELECT id FROM users LIMIT 1`
	err = db.QueryRow(query).Scan(&userID)
	if err != nil {
		return err
	}
	parsedUUID, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	doc.UserID = parsedUUID
	return nil

}

func (doc *Doc) setTitle(fileName string) error {
	doc.Title = filepath.Base(fileName)
	return nil
}

func (doc *Doc) setDirectory(fileName string) error {
	dir := filepath.Dir(fileName)
	if dir == "." {
		_loc, err := os.Getwd()
		if err != nil {
			return err
		}
		dir = _loc
	}
	doc.Directory = dir
	return nil
}

func (doc *Doc) setCreatedDate() {
	doc.CreatedDate = time.Now()
}

func (doc *Doc) setKeyword(keyword string) {
	doc.Keyword = keyword
}
