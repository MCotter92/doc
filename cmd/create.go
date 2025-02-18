/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/MCotter92/doc/docCore"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a document and tracks it with doc.",
	Long:  `Creates a document and tracks it in with doc.`,
    Args: cobra.RangeArgs(2, 2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Print("The create command requires two arguments: a file name and a keyword.")
		}
		fmt.Println("create called")
        docCore.CreateDocument(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
