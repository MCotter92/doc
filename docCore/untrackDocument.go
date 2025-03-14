package docCore
import (
    "fmt"
    "os"
	"encoding/json"
	"github.com/MCotter92/doc/utils"

)
func Untrack() {
    // unmarshal global.json

    globalData, err := os.ReadFile("~/dev/doc/data/global.json")
    if err != nil {
        fmt.Println("Cannot read global.json", err)
    }

    var store  utils.DocumentStore
    err = json.Unmarshal(globalData, &store)
    if err != nil {
        fmt.Println("Cannot unmarshal globalData", err)
    }

    // search for document and remove from struct 
    // Search()
    // marshal global.json 
    // delete document? (y/n)
}
