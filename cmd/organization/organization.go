package organization

import (
	"github.com/reverendyz/adocli/cmd/organization/set"
	"github.com/spf13/cobra"
)

var (
	OrganizationCommand = &cobra.Command{
		Use:   "organization",
		Short: "Manages organization options",
	}
)

func init() {
	OrganizationCommand.AddCommand(set.OrganizationSetCommand)
}
