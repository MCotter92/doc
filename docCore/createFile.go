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
		fmt.Printf("Could not create note struct: %w", err)
		return err
	}

	err = utils.CreateNoteFile(note.Title, note.Keyword)
	if err != nil {
		fmt.Printf("Could not create note file: %w", err)
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Could not get home dir from os package: %w", err)
		return err
	}

	dbPath := filepath.Join(homeDir, ".config/doc/doc.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Printf("Could not open db: %w", err)
		return err
	}
	defer db.Close()

	if err := utils.InsertNote(db, note); err != nil {
		fmt.Printf("Could not insert note into db: %w", err)
		return err
	}

	return nil
}
