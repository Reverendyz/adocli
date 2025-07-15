package describe

import (
	"github.com/reverendyz/adocli/pkg/teams/describe"
	"github.com/spf13/cobra"
)

var (
	GetTeamsCommand = &cobra.Command{
		Use:   "describe <organization-url> <project-id>",
		Short: "Describe team under project",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return describe.GetTeams(args[0], args[1])
		},
	}
)
