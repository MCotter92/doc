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
	notesLocation string
)

// newProfileCmd represents the newProfile command
var newProfileCmd = &cobra.Command{
	Use:   "newProfile",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		docCore.CreateProfile(editor, notesLocation)

	},
}

func init() {
	rootCmd.AddCommand(newProfileCmd)

	newProfileCmd.Flags().StringVar(&editor, "Editor", "", "Set the user's prefered editor.")
	newProfileCmd.Flags().StringVar(&notesLocation, "Notes Location", "", "Set the user's prefered notes location.")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newProfileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newProfileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
