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

	err = os.WriteFile(doc.Title, []byte(frontmatter), 0644)
	if err != nil {
		return fmt.Errorf("Could not write to config file: %w", err)
	}

	return nil
}

func (n *Doc) setPath(title, directory string) {
	n.Path = filepath.Join(title, directory)
}

func (n *Doc) setID() {
	n.Id = uuid.New()
}

func (n *Doc) GetID() uuid.UUID {
	return n.Id
}

func (n *Doc) setUserID() error {
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

	n.UserID = parsedUUID
	return nil

}

func (n *Doc) setTitle(fileName string) error {
	n.Title = filepath.Base(fileName)
	return nil
}

func (n *Doc) setDirectory(fileName string) error {
	dir := filepath.Dir(fileName)
	if dir == "." {
		_loc, err := os.Getwd()
		if err != nil {
			return err
		}
		dir = _loc
	}
	n.Directory = dir
	return nil
}

func (n *Doc) setCreatedDate() {
	n.CreatedDate = time.Now()
}

func (n *Doc) setKeyword(keyword string) {
	n.Keyword = keyword
}
