/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MCotter92/doc/utils"
	"github.com/spf13/cobra"
)

var (
	keywordFlag     string
	titleFlag       string
	locationFlag    string
	createdDateFlag string
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		criteria := utils.SearchCriteria{
			Keyword:     keywordFlag,
			Title:       titleFlag,
			Location:    locationFlag,
			CreatedDate: createdDateFlag,
		}

		homedir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error getting home director: %v\n", err)
		}
		dbPath := filepath.Join(homedir, ".config", "doc", "doc.db")
		//TODO: figure out my custom NewDatabase mess.
		db, err := utils.NewDatabase(dbPath)
		if err != nil {
			fmt.Printf("could not open db: %v", err)
		}
		defer db.DB.Close()

		notes, err := db.Search(criteria)
		if err != nil {
			fmt.Printf("search faild: %v\n", err)
		}

		if len(notes) == 0 {
			fmt.Println("No documents found matching your criteria.")
			return
		}

		fmt.Printf("Found %d documents(s):\n\n", len(notes))
		for _, doc := range notes {
			fmt.Printf("ID: %d\n", doc.Id)
			fmt.Printf("Keyword: %s\n", doc.Keyword)
			fmt.Printf("Title: %s\n", doc.Title)
			fmt.Printf("Location: %s\n", doc.Location)
			fmt.Printf("CreatedDate: %s\n", doc.CreatedDate.String())
			fmt.Println("-------------------------------")
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	openCmd.Flags().StringVarP(&keywordFlag, "keyword", "k", "", "Search by keyword")
	openCmd.Flags().StringVarP(&titleFlag, "title", "t", "", "Search by title")
	openCmd.Flags().StringVarP(&locationFlag, "location", "l", "", "Search by location")
	openCmd.Flags().StringVarP(&createdDateFlag, "created", "c", "", "Search by created date")
}
