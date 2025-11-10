/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/MCotter92/doc/docCore"
	"github.com/spf13/cobra"
)

var (
	editor        string
	userName      string
	notesLocation string
	showConfig    bool
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: add viper file watcher to config file to watch for changes there
		hasUpdates := editor != "" || userName != "" || notesLocation != ""

		if hasUpdates {
			// updateReq := docCore.ConfigUpdateRequest{
			// 	Editor:        editor,
			// 	UserName:      userName,
			// 	NotesLocation: notesLocation,
			// }
			// return docCore.UpdateUserConfig(updateReq)
		}

		return docCore.ShowUserConfig(showConfig)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().StringVarP(&editor, "editor", "e", "", "Set the default editor (e.g., nvim, code, emacs)")
	configCmd.Flags().StringVarP(&userName, "username", "u", "", "Set the username")
	configCmd.Flags().StringVarP(&notesLocation, "notes-location", "n", "", "Set the notes location directory")
	configCmd.Flags().BoolVar(&showConfig, "show", false, "Show detailed configuration including internal paths")
}
