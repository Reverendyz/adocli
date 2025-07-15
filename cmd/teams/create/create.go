package create

import (
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/reverendyz/adocli/pkg/teams/create"
	"github.com/spf13/cobra"
)

var (
	CreateTeamCommand = &cobra.Command{
		Use:   "create <team-name> <project-name>",
		Short: "Creates a new team in the specified project",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return create.CreateTeam(&core.WebApiTeam{
				Name:        &args[0],
				ProjectName: &args[1],
			})
		},
	}
)
