package list

import (
	"github.com/reverendyz/adocli/config"
	"github.com/reverendyz/adocli/pkg/projects/list"
	"github.com/spf13/cobra"
)

var (
	ProjectsListCommand = &cobra.Command{
		Use: "list",
		Run: func(cmd *cobra.Command, args []string) {
			list.ProjectsList(config.OrganizationUrl)
		},
	}
)
