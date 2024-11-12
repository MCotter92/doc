/*
Copyright © 2024 NAME HERE <maccotter11@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/MCotter92/doc/utils"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// makeCmd represents the make command
var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "Makes either a txt or a markdown file to be tracked by doc.",
	Long:  "Makes a file with the specified parameters: Title, Extension, Locaition, and a list of keywords.",
	Run: func(cmd *cobra.Command, args []string) {
		//TODO Right now there are no optional params and so the only param currently is title. Fix this? Pointer trick from Jake?
		makeDocument(args[0])
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)

}

func makeDocument(Title string) utils.Document {

	newDocument := utils.Document{
		UUID:        uuid.NewString(),
		Title:       filepath.Base(Title),
		Extension:   filepath.Ext(Title),
		Location:    filepath.Dir(Title),
		CreatedDate: time.Now().Format("2006-01-02"),
	}

	docStruct, err := os.Create(newDocument.Title)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
	}
	defer docStruct.Close()

	docJson, err := json.MarshalIndent(newDocument, "", "\t")
	if err != nil {
		fmt.Printf("Error marshalling struct: %v\n", err)
	}

	globalJson, err := os.OpenFile("data/global.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := globalJson.Write([]byte(docJson)); err != nil {
		globalJson.Close()
		log.Fatal(err)
	}

	globalJson.Close()

	fmt.Println(newDocument)
	return newDocument
}
