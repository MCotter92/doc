package docCore

import (
	"fmt"
	"os"

	"github.com/MCotter92/doc/utils"
)

func Search(searchCriteria utils.SearchCriteria) ([]utils.Doc, error) {

	db, err := utils.NewDatabase()
	if err != nil {
		return nil, fmt.Errorf("Could not open db: %s", err)
	}
	defer db.Close()

	res, err := db.Search(searchCriteria)
	if err != nil {
		return nil, fmt.Errorf("Search failed: %s", err)
	}

	if len(res) == 0 {
		fmt.Println("No documents found matching your criteria.")
		os.Exit(0)
	}

	return res, nil
}
