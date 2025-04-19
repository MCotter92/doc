package docCore

import (
	"database/sql"
	"github.com/MCotter92/doc/utils"
	_ "github.com/mattn/go-sqlite3"
)

func TrackDocument(db *sql.DB, title, keyword string) {

	// fill out child struct
	pDoc := &utils.Document{}
	pDoc.SetID()
	pDoc.SetTitle(title)
	pDoc.SetExtension(title)
	pDoc.SetLocation(title)
	pDoc.SetFullName()
	pDoc.SetInode()
	pDoc.SetCreatedDate()
	pDoc.SetKeyword(keyword)

}
