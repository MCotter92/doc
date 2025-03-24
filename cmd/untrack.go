/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	docCore "github.com/MCotter92/doc/docCore"
	"github.com/spf13/cobra"
)

// untrackCmd represents the untrack command
var untrackCmd = &cobra.Command{
	Use:   "untrack",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
        docCore.Untrack(id, title, extension, location, createdDate, keyword)
		fmt.Println("untrack called")
	},
}

func init() {
	rootCmd.AddCommand(untrackCmd)

    untrackCmd.Flags().StringVar(&id, "id", "", "Search doc by keyword.")
    untrackCmd.Flags().StringVar(&title, "title", "", "Search doc by title.")
    untrackCmd.Flags().StringVar(&extension, "exension", "", "Search doc by extension.")
    untrackCmd.Flags().StringVar(&location, "location", "", "Search doc by location.")
    untrackCmd.Flags().StringVar(&createdDate, "createdDate", "", "Search doc by created date.")
    untrackCmd.Flags().StringVar(&keyword, "keyword", "", "Search doc by keyword.")
}
