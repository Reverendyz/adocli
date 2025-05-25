/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/reverendyz/adocli/cmd/item"
	"github.com/reverendyz/adocli/cmd/projects"
	"github.com/reverendyz/adocli/cmd/teams"
	"github.com/spf13/cobra"
)

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
	rootCmd.AddCommand(teams.TeamCmd)

	rootCmd.Flags().StringP("organization", "o", "", "organization string to execute the commands")
}
