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
		Short: "Handles Azure Devops teams",
	}
)

func init() {
	TeamCmd.AddCommand(describe.GetTeamsCommand)
	TeamCmd.AddCommand(create.CreateTeamCommand)
	TeamCmd.AddCommand(delete.DeleteTeamCommand)
	TeamCmd.Flags().StringP("project-id", "p", "", "project id")
}
