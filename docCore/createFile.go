package docCore

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/MCotter92/doc/utils"
)

func CreateFile(title, keyword string) error {
	note, err := utils.NewNote(title, keyword)
	if err != nil {
		return fmt.Errorf("Could not create note struct: %w", err)
	}

	err = utils.CreateNoteFile(note.Title, note.Keyword)
	if err != nil {
		return fmt.Errorf("Could not create note file: %w", err)
	}

	homeDir, _ := os.UserHomeDir()
	dbPath := filepath.Join(homeDir, ".config/doc/doc.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("Could not open db: %w", err)
	}
	defer db.Close()
	if err := utils.InsertNote(db, note); err != nil {
		return fmt.Errorf("Could not insert note into db: %w", err)
	}

	return nil

}
