package docCore

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/MCotter92/doc/utils"
)

func Search(id, title, extension, location, createdDate,  keyword string) []string{
    globalData, err := os.ReadFile("data/global.json")
    if err != nil {
        fmt.Println("Cannot read global.json", err)
    }

    var store utils.DocumentStore
    err = json.Unmarshal(globalData, &store)
    if err != nil {
        fmt.Println("Cannon unmarshal globalData", err)
    }

    var docs []string
    for _, doc := range store.Documents {
        if id != "" && doc.Id.String() != id {
            continue
        }

        if title != "" && !strings.Contains(doc.Title, title) {
            continue
        }

        if extension != "" && doc.Extension != extension {
            continue
        }

        if location != "" && !strings.Contains(doc.Location, location) {
            continue
        }

        if createdDate != "" && doc.CreatedDate.String() != createdDate {
            continue
        }


        if keyword != "" && !strings.Contains(doc.Keyword, keyword) {
            continue
        }

        docs = append(docs, doc.Location)

    }

    return docs
    
}
