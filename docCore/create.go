package docCore

import (
	"fmt"

	"github.com/MCotter92/doc/utils"
)

func CreateDoc(title, keyword string, db *utils.Database) error {
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

	if err := db.InsertDoc(doc); err != nil {
		fmt.Printf("Could not insert note into db: %s", err)
		return err
	}

	return nil
}
