package utils

import (
	"fmt"
)

func SetFrontmatter(note *Note) string {
	frontmatter := fmt.Sprintf(`---
	Id         %s 
	User_id    %s 
	Title      %s 
	Extension  %s 
	Location   %s 
	CreatedDate %s 
	Keyword    %s 
	---`,
		note.Id, note.User_id, note.Title, note.Extension, note.Location, note.CreatedDate, note.Keyword)

	return frontmatter

}

// TODO: finish this
// func ParseFrontmatter(filePath string) (FrontMatter, error)
