package get

import (
	"github.com/reverendyz/adocli/pkg"
	"github.com/spf13/cobra"
)

var (
	GetProjectsCommand = &cobra.Command{
		Use: "get",
		Run: func(cmd *cobra.Command, args []string) {
			pkg.TestFunc(args[0])
		},
	}
)
