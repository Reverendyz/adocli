package teams

import (
	"github.com/reverendyz/adocli/cmd/teams/create"
	"github.com/reverendyz/adocli/cmd/teams/delete"
	"github.com/reverendyz/adocli/cmd/teams/describe"
	"github.com/spf13/cobra"
)

var (
	TeamCmd = &cobra.Command{
		Use:   "teams",
		Short: "Manage Azure DevOps teams",
		Long:  "Manage Azure DevOps teams including create, delete, and describe operations.",
	}
)

func init() {
	TeamCmd.AddCommand(describe.GetTeamsCommand)
	TeamCmd.AddCommand(create.CreateTeamCommand)
	TeamCmd.AddCommand(delete.DeleteTeamCommand)
	TeamCmd.Flags().StringP("project-id", "p", "", "project id")
}
