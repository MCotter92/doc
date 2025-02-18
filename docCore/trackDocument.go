package docCore

import 
(
	"encoding/json"
	"fmt"
	"os"
	"time"
	"path/filepath"

	"github.com/MCotter92/doc/utils"
	"github.com/google/uuid"
)


func TrackDocument(fileName ,keyword string)  {

        
    // read data/global.json for use 
    globalData, err := os.ReadFile("data/global.json")
    if err != nil {
        fmt.Println("Cannot read global.json", err)
    }

    // unmarshal global.json and store into parent struct
    var store  utils.DocumentStore
    err = json.Unmarshal(globalData, &store)
    if err != nil {
        fmt.Println("Cannot unmarshal globalData", err)
    }

    loc := filepath.Dir(fileName)
    if loc == "."{
        _loc, err := os.Getwd()
        if err != nil {
            fmt.Printf("Cannot retrieve file location: %v\n", err)
        }
        loc = _loc
    }

    stats,_:=os.Stat(fileName)

    ext := filepath.Ext(fileName)
    id := uuid.New()

    // fill out child struct
    d := utils.Document{
        Id:               id,
        Title:            stats.Name(),
        Extension:        ext,
        Location:         loc,
        CreatedDate:      time.Now(),
        LastModifiedDate: stats.ModTime(),
        Keyword:         keyword,
        
    }

    store.Documents = append(store.Documents, d)

    b, err := json.MarshalIndent(store, "", " ")
    if err != nil {
        fmt.Println("Cannot marshal store", err)
    }

    err  = os.WriteFile("data/global.json", b, 0644)
    if err != nil {
        fmt.Println("Cannot write to data/global.json", err)
    }

}
