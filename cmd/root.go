/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/reverendyz/adocli/cmd/item"
	"github.com/reverendyz/adocli/cmd/projects"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "adocli",
	Short: "A command line interface to manage Azure DevOps items",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(item.ItemCmd)
	rootCmd.AddCommand(projects.ProjectsCmd)
}
