/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/reverendyz/adocli/cmd/item"
	"github.com/reverendyz/adocli/cmd/organization"
	"github.com/reverendyz/adocli/cmd/projects"
	"github.com/reverendyz/adocli/cmd/teams"
	"github.com/reverendyz/adocli/common"
	"github.com/reverendyz/adocli/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "adocli",
	Short: "A command line interface to manage Azure DevOps items",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		config.ProjectId, err = common.SetOrLoadFlag(cmd, "project", "projectId")
		if err != nil {
			fmt.Println("No projectId set")
			os.Exit(1)
		}
		config.OrganizationUrl, err = common.SetOrLoadFlag(cmd, "organization", "organizationUrl")
		if err != nil {
			fmt.Println("No organizationUrl set")
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("organization", "o", "", "organization string to be used in a single run")
	rootCmd.PersistentFlags().StringP("project", "p", "", "project id string to be used in a single run")

	rootCmd.AddCommand(item.ItemCmd)
	rootCmd.AddCommand(projects.ProjectsCmd)
	rootCmd.AddCommand(teams.TeamCmd)
	rootCmd.AddCommand(organization.OrganizationCommand)

}
