package utils

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"time"
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

// TODO: Finish these
// func SyncNoteWithDB(db *Database, filePath string, frontmatter FrontMatter) error

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
	// get user_id from users table
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

func (n *Note) setTitle(fileName string) error {
	stats, err := os.Stat(fileName)
	if err != nil {
		return err
	}
	n.Title = stats.Name()
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
