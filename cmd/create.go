/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a document and tracks it with doc.",
	Long:  `Creates a document and tracks it in with doc.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVar(&title, "title", "", "Search doc by title.")
	createCmd.Flags().StringVar(&keyword, "keyword", "", "Search doc by keyword.")

}
