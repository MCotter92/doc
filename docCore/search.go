package docCore

import (
	"fmt"
	"os"

	"github.com/MCotter92/doc/utils"
	_ "github.com/mattn/go-sqlite3"
)

func Search(searchCriteria utils.DocumentSearchCriteria) ([]utils.Doc, *utils.Database, error) {

	db, err := utils.NewDatabase()
	if err != nil {
		return nil, nil, fmt.Errorf("Could not open db: %s", err)
	}

	res, err := db.SearchDocumentsTable(searchCriteria)
	if err != nil {
		db.Close()
		return nil, nil, fmt.Errorf("Search failed: %s", err)
	}

	if len(res) == 0 {
		fmt.Println("No documents found matching your criteria.")
		db.Close()
		os.Exit(0)
	}

	return res, db, nil
}
