package docCore

import 
(
	"encoding/json"
	"fmt"
	"os"
	

	"github.com/MCotter92/doc/utils"
)


func TrackDocument(title, keyword string)  {

        
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

    // fill out child struct
    pDoc := &utils.Document{}
    pDoc.SetID()
    pDoc.SetTitle(title)
    pDoc.SetExtension(title)
    pDoc.SetLocation(title)
    pDoc.SetCreatedDate()
    pDoc.SetKeyword(keyword)

    // append new child struct to docs list in parent struct
    store.Documents = append(store.Documents, *pDoc)

    // marshal parent struct into json
    b, err := json.MarshalIndent(store, "", " ")
    if err != nil {
        fmt.Println("Cannot marshal store", err)
    }

    // write parent struct to global.json 
    err  = os.WriteFile("data/global.json", b, 0644)
    if err != nil {
        fmt.Println("Cannot write to ~/dev/doc/data/global.json", err)
    }

}
