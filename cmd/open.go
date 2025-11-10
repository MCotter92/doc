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

var (
	idFlag          string
	userIdFlag      string
	directoryFlag   string
	titleFlag       string
	pathFlag        string
	createdDateFlag string
	keywordFlag     string
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

		searchCriteria := utils.SearchCriteria{
			Id:          idFlag,
			UserID:      userIdFlag,
			Directory:   directoryFlag,
			Title:       titleFlag,
			Path:        pathFlag,
			CreatedDate: createdDateFlag,
			Keyword:     keywordFlag,
		}

		searchRes, _, err := docCore.Search(searchCriteria)
		if err != nil {
			fmt.Println(err)
		}

		err = docCore.Open(searchRes)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	openCmd.Flags().StringVarP(&idFlag, "Id", "id", "", "Search by doc id.")
	openCmd.Flags().StringVarP(&userIdFlag, "userId", "uid", "", "Search by doc user id.")
	openCmd.Flags().StringVarP(&directoryFlag, "directory", "d", "", "Search by a doc's directory.")
	openCmd.Flags().StringVarP(&titleFlag, "title", "t", "", "Search by a doc's title.")
	openCmd.Flags().StringVarP(&pathFlag, "path", "f", "", "Search by a doc's full path.")
	openCmd.Flags().StringVarP(&createdDateFlag, "created", "c", "", "Search by created date")
	openCmd.Flags().StringVarP(&keywordFlag, "keyword", "k", "", "Search by keyword")
}
