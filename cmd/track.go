/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	// "github.com/MCotter92/doc/utils"
	// "github.com/djherbis/times"
	// "github.com/google/uuid"
	"github.com/spf13/cobra"
	// "os"
	// "path/filepath"
)

// trackCmd represents the track command
var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Tracks a document.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("track called")
	},
}

func init() {
	rootCmd.AddCommand(trackCmd)

}

// func trackDocument(item string) utils.Document {
// 	//TODO trackDocument takes in a string and checks if the string represents a document or a directory. If it is a directroy doc should
// 	//ask the user if they want to track everything in the repository or not. If yes, track all documents in that directory.
//
// 	//TODO add a flag -r that runs this recursively so that it trackes everything below the cwd including nested dirs and stuff.
//
// 	Location := filepath.Dir(item)
// 	if Location == "." {
// 		var err error
// 		Location, err = os.Getwd()
// 		if err != nil {
// 			fmt.Printf("Cannot retrieve file location: %v\n", err)
// 		}
// 	}
//
// 	info, err := os.Stat(item)
// 	if err != nil {
// 		fmt.Printf("Cannot retrieve file stats: %v\n (y/n)", err)
// 	}
//
// 	// if info.IsDir() {
// 	// 	var userInput string
// 	// 	fmt.Print("Do you want doc to track all files within %v?", item)
// 	// 	fmt.Scanln(&userInput)
// 	// 	if userInput == "y" || userInput == "Y" {
// 	// 		filepath.WalkDir(item, makeDocument())
// 	// 		// TODO walk the directory and convert each FILE to my struct then add to global.json.
// 	// 		// Need to refactor the tracking functionality to utils?
// 	//
// 	// 	} else {
// 	// 		fmt.Println("doc track cancelled.")
// 	// 	}
// 	// }
// 	//
// 	// // Create a doc struct
// 	// return newDocument
//
// }
