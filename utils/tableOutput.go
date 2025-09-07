package utils

import (
	"os"
	"strconv"

	"github.com/aquasecurity/table"
)

func TableOutput(searchRes []Doc) *table.Table {
	t := table.New(os.Stdout)

	t.SetHeaders("Row Num", "ID", "Keyword", "Path", "Created Date")
	for i, doc := range searchRes {
		t.AddRow(strconv.Itoa(i), doc.Id.String(), doc.Keyword, doc.Path, doc.CreatedDate.String())
	}

	t.Render()

	return t
}
