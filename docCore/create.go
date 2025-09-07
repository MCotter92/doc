package docCore

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/MCotter92/doc/utils"
)

func CreateFile(title, keyword string) error {
	doc, err := utils.NewDoc(title, keyword)
	if err != nil {
		fmt.Printf("Could not create note struct: %s", err)
		return err
	}

	err = utils.CreateDocFile(doc.Title, doc.Keyword)
	if err != nil {
		fmt.Printf("Could not create note file: %s", err)
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Could not get home dir from os package: %s", err)
		return err
	}

	dbPath := filepath.Join(homeDir, ".config/doc/doc.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Printf("Could not open db: %s", err)
		return err
	}
	defer db.Close()

	if err := utils.InsertNote(db, doc); err != nil {
		fmt.Printf("Could not insert note into db: %s", err)
		return err
	}

	return nil
}
