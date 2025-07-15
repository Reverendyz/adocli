package set

import (
	"github.com/reverendyz/adocli/pkg/organization/set"
	"github.com/spf13/cobra"
)

var (
	OrganizationSetCommand = &cobra.Command{
		Use:   "set <organization-url>",
		Short: "sets organization URL in cache for further usage. For single run, use -o <organization URL> option.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return set.SetOrganizationInConfigFile(args[0])
		},
	}
)
