package list

import (
	"github.com/reverendyz/adocli/config"
	"github.com/reverendyz/adocli/pkg/projects/list"
	"github.com/spf13/cobra"
)

var (
	ProjectsListCommand = &cobra.Command{
		Use:   "list",
		Short: "List all projects in the organization",
		RunE: func(cmd *cobra.Command, args []string) error {
			return list.ProjectsList(config.OrganizationUrl)
		},
	}
)
