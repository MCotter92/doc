/*
Copyright © 2024 NAME HERE <maccotter11@gmail.com>
*/
// package cmd
//
// import (
// 	"encoding/json"
// 	"fmt"
// 	"github.com/MCotter92/doc/utils"
// 	"github.com/djherbis/times"
// 	"github.com/google/uuid"
// 	"github.com/spf13/cobra"
// 	"log"
// 	"os"
// 	"path/filepath"
// )
//
// // makeCmd represents the make command
// var makeCmd = &cobra.Command{
// 	Use:   "make",
// 	Short: "Makes either a txt or a markdown file to be tracked by doc.",
// 	Long:  "Makes a file with the specified parameters: Title, Extension, Locaition, and a list of keywords.",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		//TODO Right now there are no optional params and so the only param currently is title. Fix this? Pointer trick from Jake?
// 		makeDocument(args[0])
//
// 	},
// }
//
// func init() {
// 	rootCmd.AddCommand(makeCmd)
// 	// TODO define flags.
//
// }
//
// func makeDocument(Title string) utils.Document {
//
// 	Location := filepath.Dir(Title)
// 	if Location == "." {
// 		var err error
// 		Location, err = os.Getwd()
// 		if err != nil {
// 			fmt.Printf("Cannot retrieve file location: %v\n", err)
// 		}
// 	}
//
// 	times, err := times.Stat(Title)
// 	if err != nil {
// 		fmt.Printf("Cannot retrieve file times: %v\n", err)
// 	}
//
//
// 	// Create a file using the struct
// 	docStructFile, err := os.Create(newDocument.Title)
// 	if err != nil {
// 		fmt.Printf("Error creating file: %v\n", err)
// 	}
// 	defer docStructFile.Close()
//
// 	// Convert the struct into Json
// 	docJson, err := json.MarshalIndent(newDocument, "", "\t")
// 	if err != nil {
// 		fmt.Printf("Error marshalling struct: %v\n", err)
// 	}
//
// 	// Open the Json file
// 	globalJson, err := os.OpenFile("data/global.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	// write the struct turned json to global.json
// 	if _, err := globalJson.Write([]byte(docJson)); err != nil {
// 		globalJson.Close()
// 		log.Fatal(err)
// 	}
//
// 	globalJson.Close()
//
// 	fmt.Println(newDocument)
// 	return newDocument
// }
