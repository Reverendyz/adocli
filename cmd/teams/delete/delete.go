package delete

import (
	"fmt"

	"github.com/reverendyz/adocli/pkg/teams/delete"
	"github.com/spf13/cobra"
)

var (
	DeleteTeamCommand = &cobra.Command{
		Use:   "delete",
		Short: "Command to create new team from",
		Run: func(cmd *cobra.Command, args []string) {
			teamName := args[0]
			fmt.Println(teamName)
			err := delete.DeleteTeam(teamName)
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		},
	}
)
