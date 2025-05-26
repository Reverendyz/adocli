package set

import (
	"github.com/reverendyz/adocli/pkg/organization/set"
	"github.com/spf13/cobra"
)

var (
	OrganizationSetCommand = &cobra.Command{
		Use:   "set",
		Short: "sets organization URL in cache for further usage. For single run, use -o <organization URL> option.",
		Run: func(cmd *cobra.Command, args []string) {
			set.SetOrganizationInConfigFile(args[0])
		},
	}
)
