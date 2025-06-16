package docCore

import (
	"fmt"
	"os"

	"github.com/MCotter92/doc/utils"
)

func CreateFile(title, keyword string) error {
	note, err := utils.NewNote(title, keyword)
	if err != nil {
		return fmt.Errorf("Could not create note: %w", err)
	}

	frontmatter := utils.SetFrontmatter(note)

	// Write config file
	err = os.WriteFile(note.Title, []byte(frontmatter), 0644)
	if err != nil {
		return fmt.Errorf("Could not write to config file: %w", err)
	}

	return nil

}
