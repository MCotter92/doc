/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/MCotter92/doc/docCore"
	"github.com/spf13/cobra"
)

var (
	path    string
	keyword string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		docCore.CreateFile(path, keyword)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&path, "path", "p", "", "Path for a new note.")
	createCmd.Flags().StringVarP(&keyword, "keyword", "k", "", "Keyword for a new note.")
}
