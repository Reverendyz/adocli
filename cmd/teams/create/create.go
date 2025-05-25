package create

import (
	"fmt"

	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/reverendyz/adocli/pkg/teams/create"
	"github.com/spf13/cobra"
)

var (
	CreateTeamCommand = &cobra.Command{
		Use:   "create",
		Short: "Command to create new team from",
		Run: func(cmd *cobra.Command, args []string) {
			err := create.CreateTeam(&core.WebApiTeam{
				Name:        &args[0],
				ProjectName: &args[1],
			})
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		},
	}
)
