package docCore

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/MCotter92/doc/utils"
)
func Untrack(id, title, extension, location, createdDate,  keyword string) {
    // unmarshal global.json

    var userInput string
    fmt.Println("Would you like to delete the document as well as untrack it?")
    fmt.Scanln(&userInput)

    globalData, err := os.ReadFile("data/global.json")
    if err != nil {
        fmt.Println("Cannot read global.json", err)
    }

    var store  utils.DocumentStore
    err = json.Unmarshal(globalData, &store)
    if err != nil {
        fmt.Println("Cannot unmarshal globalData", err)
    }

    // search for document and remove from struct 
    var remainingDocuments []utils.Document
    for _, doc := range store.Documents {
        if !((id != "" && doc.Id.String() == id) ||
        (title != "" && strings.Contains(doc.Title, title)) ||
        (extension != "" && doc.Extension == extension) ||
        (location != "" && strings.Contains(doc.Location, location)) ||
        (createdDate != "" && doc.CreatedDate.String() == createdDate) ||
        (keyword != "" && strings.Contains(doc.Keyword, keyword))) {
            // If it doesn't meet any filter criteria, keep the document in the new slice
            remainingDocuments = append(remainingDocuments, doc)
        }
    }

    store.Documents = remainingDocuments

    b, err := json.MarshalIndent(store, "", " ")
    if err != nil {
        fmt.Println("Cannot marshal store", err)
    }

    err  = os.WriteFile("data/global.json", b, 0644)
    if err != nil {
        fmt.Println("Cannot write to ~/dev/doc/data/global.json", err)
    }
}
