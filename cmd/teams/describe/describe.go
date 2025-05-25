package describe

import (
	"github.com/reverendyz/adocli/pkg/teams/describe"
	"github.com/spf13/cobra"
)

var (
	GetTeamsCommand = &cobra.Command{
		Use:   "describe",
		Short: "Describe team under project",
		Run: func(cmd *cobra.Command, args []string) {
			organizationUrl := args[0]
			projectId := args[1]
			describe.GetTeams(organizationUrl, projectId)
		},
	}
)
