package delete

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	CreateTeamCommand = &cobra.Command{
		Use:   "create",
		Short: "Command to create new team from",
		Run: func(cmd *cobra.Command, args []string) {
			err := delete.DeleteTeam(os.Args[0])
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		},
	}
)
