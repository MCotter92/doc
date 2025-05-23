/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	id          string
	title       string
	extension   string
	location    string
	createdDate string
	keyword     string
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for tracked docs.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// docs := docCore.Search(id, title, extension, location, createdDate, keyword)
		// for _,v := range docs{
		//     fmt.Println(v)
		// }
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringVar(&id, "id", "", "Search doc by keyword.")
	searchCmd.Flags().StringVar(&title, "title", "", "Search doc by title.")
	searchCmd.Flags().StringVar(&extension, "exension", "", "Search doc by extension.")
	searchCmd.Flags().StringVar(&location, "location", "", "Search doc by location.")
	searchCmd.Flags().StringVar(&createdDate, "createdDate", "", "Search doc by created date.")
	searchCmd.Flags().StringVar(&keyword, "keyword", "", "Search doc by keyword.")
}
