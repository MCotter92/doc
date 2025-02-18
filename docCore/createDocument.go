package docCore

import (
	"fmt"
	"os"
	"strings"
)



func CreateDocument(fileName  string, keyword string )  {

    // returns the full path to the current dir
    if FileExists(fileName) {
        var userInput string
        fmt.Println("This file already exists. Would you like to track it? (y/n)")
        fmt.Scanln(&userInput)
        if strings.ToLower(userInput) == "y" {
            TrackDocument(fileName, keyword)
        }
        if strings.ToLower(userInput) == "n" {
            fmt.Println("File was not tracked.")
        }
    } else {
        var userInput2 string 
        fmt.Println("This file does not exists. Would you like to creat it and then track it? (y/n)")
        fmt.Scanln(&userInput2)
        if strings.ToLower(userInput2) == "y" {
            file, err := os.Create(fileName)
            if err != nil {
                fmt.Println("Could not create file:", err)
            }
            defer file.Close()

            TrackDocument(fileName, keyword)
        }
        if strings.ToLower(userInput2) == "n" {
            fmt.Println("File was not created or tracked.")
        }
    }
}

