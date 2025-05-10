package get

import (
	"github.com/reverendyz/adocli/pkg/teams/get"
	"github.com/spf13/cobra"
)

var (
	GetTeamsCommand = &cobra.Command{
		Use: "get",
		Run: func(cmd *cobra.Command, args []string) {
			organizationUrl := args[0]
			projectId := args[1]
			get.GetTeams(organizationUrl, projectId)
		},
	}
)
