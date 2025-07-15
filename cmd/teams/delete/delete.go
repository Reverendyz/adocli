package delete

import (
	"github.com/reverendyz/adocli/pkg/teams/delete"
	"github.com/spf13/cobra"
)

var (
	DeleteTeamCommand = &cobra.Command{
		Use:   "delete <team-name>",
		Short: "Deletes a team by name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return delete.DeleteTeam(args[0])
		},
	}
)
