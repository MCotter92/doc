package utils

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type Note struct {
	Id          uuid.UUID
	User_id     uuid.UUID
	Title       string
	Extension   string
	Location    string
	CreatedDate time.Time
	Keyword     string
}

func CreateNoteFile(title, keyword string) error {

	note, err := NewNote(title, keyword)
	if err != nil {
		return fmt.Errorf("Could not create note: %w", err)
	}

	frontmatter := SetFrontmatter(note)

	// Write config file
	err = os.WriteFile(note.Title, []byte(frontmatter), 0644)
	if err != nil {
		return fmt.Errorf("Could not write to config file: %w", err)
	}

	return nil
}

func NewNote(title, keyword string) (*Note, error) {
	doc := &Note{}

	doc.setID()
	doc.setUserID()
	if err := doc.setTitle(title); err != nil {
		return nil, err
	}
	doc.setExtension(title)
	if err := doc.setLocation(title); err != nil {
		return nil, err
	}
	doc.setCreatedDate()
	doc.setKeyword(keyword)

	return doc, nil

}

func (n *Note) setID() {
	n.Id = uuid.New()
}

func (n *Note) GetID() uuid.UUID {
	return n.Id
}

func (n *Note) setUserID() error {
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
	parsedUUDI, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	n.User_id = parsedUUDI
	return nil

}

func (n *Note) setTitle(fileName string) error {
	n.Title = filepath.Base(fileName)
	return nil
}

func (n *Note) setExtension(fileName string) {
	ext := filepath.Ext(fileName)
	n.Extension = ext
}

func (n *Note) setLocation(fileName string) error {
	loc := filepath.Dir(fileName)
	if loc == "." {
		_loc, err := os.Getwd()
		if err != nil {
			return err
		}
		loc = _loc
	}
	n.Location = loc
	return nil
}

func (n *Note) setCreatedDate() {
	n.CreatedDate = time.Now()
}

func (n *Note) setKeyword(keyword string) {
	n.Keyword = keyword
}
