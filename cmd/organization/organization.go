package organization

import (
	"github.com/reverendyz/adocli/cmd/organization/set"
	"github.com/spf13/cobra"
)

var (
	OrganizationCommand = &cobra.Command{
		Use:   "organization",
		Short: "Manage organization configuration",
		Long:  "Manage organization options and configuration for Azure DevOps.",
	}
)

func init() {
	OrganizationCommand.AddCommand(set.OrganizationSetCommand)
}
