package utils

import (
	"fmt"
)

func SetFrontmatter(doc *Doc) string {
	frontmatter := fmt.Sprintf(`---
	Id         %s 
	UserID    %s 
	Title      %s 
	Path   %s 
	CreatedDate %s 
	Keyword    %s 
	---`,
		doc.Id, doc.UserID, doc.Title, doc.Path, doc.CreatedDate, doc.Keyword)

	return frontmatter

}
