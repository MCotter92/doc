/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/MCotter92/doc/docCore"
	"github.com/MCotter92/doc/utils"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		searchCriteria := utils.SearchCriteria{
			Keyword:     keywordFlag,
			Title:       titleFlag,
			Path:        pathFlag,
			CreatedDate: createdDateFlag,
		}

		searchRes, db, err := docCore.Search(searchCriteria)
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()

		err = docCore.Delete(searchRes, db)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&keywordFlag, "keyword", "k", "", "Search by keyword")
	deleteCmd.Flags().StringVarP(&titleFlag, "title", "t", "", "Search by title")
	deleteCmd.Flags().StringVarP(&pathFlag, "location", "l", "", "Search by location")
	deleteCmd.Flags().StringVarP(&createdDateFlag, "created", "c", "", "Search by created date")
}
