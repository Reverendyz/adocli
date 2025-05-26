package set

import (
	"fmt"

	"github.com/reverendyz/adocli/pkg/projects/set"
	"github.com/spf13/cobra"
)

var (
	ProjectSetCommand = &cobra.Command{
		Use:   "set",
		Short: "sets project id to the cache",
		Run: func(cmd *cobra.Command, args []string) {
			if err := set.SetProjectInConfigFile(args[0]); err != nil {
				fmt.Printf("Error %v\n", err)
			}
		},
	}
)
