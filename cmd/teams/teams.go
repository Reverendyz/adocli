package teams

import (
	"github.com/reverendyz/adocli/cmd/teams/get"
	"github.com/spf13/cobra"
)

var (
	TeamCmd = &cobra.Command{
		Use:   "teams",
		Short: "handles projects commands",
	}
)

func init() {
	TeamCmd.AddCommand(get.GetTeamsCommand)
}
