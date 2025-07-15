package set

import (
	"github.com/reverendyz/adocli/pkg/projects/set"
	"github.com/spf13/cobra"
)

var (
	ProjectSetCommand = &cobra.Command{
		Use:   "set <project-id>",
		Short: "sets project id to the cache",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return set.SetProjectInConfigFile(args[0])
		},
	}
)
