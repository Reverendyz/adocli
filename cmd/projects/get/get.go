package get

import (
	"github.com/reverendyz/adocli/pkg/projects/get"
	"github.com/spf13/cobra"
)

var (
	GetProjectsCommand = &cobra.Command{
		Use: "get",
		Run: func(cmd *cobra.Command, args []string) {
			get.GetProjects(args[0])
		},
	}
)
