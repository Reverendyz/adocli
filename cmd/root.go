package cmd

import (
	"github.com/reverendyz/adocli/cmd/item"
	"github.com/reverendyz/adocli/cmd/organization"
	"github.com/reverendyz/adocli/cmd/projects"
	"github.com/reverendyz/adocli/cmd/teams"
	"github.com/reverendyz/adocli/common"
	"github.com/reverendyz/adocli/config"
	"github.com/reverendyz/adocli/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "adocli",
	Short: "A command line interface to manage Azure DevOps items",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		debug, _ := cmd.Flags().GetBool("debug")
		if err := logger.InitLogger(debug); err != nil {
			logger.FatalF("Failed to initialize logger: %v", err)
		}
		defer logger.Sync()

		var err error
		config.ProjectId, err = common.SetOrLoadFlag(cmd, "project", "projectId")
		if err != nil {
			logger.Fatal("No projectId set")
		}
		config.OrganizationUrl, err = common.SetOrLoadFlag(cmd, "organization", "organizationUrl")
		if err != nil {
			logger.Fatal("No organizationUrl set")
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		logger.Sync()
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("organization", "o", "", "organization string to be used in a single run")
	rootCmd.PersistentFlags().StringP("project", "p", "", "project id string to be used in a single run")
	rootCmd.PersistentFlags().Bool("debug", false, "enable debug logging")

	rootCmd.AddCommand(item.ItemCmd)
	rootCmd.AddCommand(projects.ProjectsCmd)
	rootCmd.AddCommand(teams.TeamCmd)
	rootCmd.AddCommand(organization.OrganizationCommand)

}
