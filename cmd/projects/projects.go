package projects

import (
	"github.com/reverendyz/adocli/cmd/projects/list"
	"github.com/reverendyz/adocli/cmd/projects/set"
	"github.com/spf13/cobra"
)

var (
	ProjectsCmd = &cobra.Command{
		Use:   "projects",
		Short: "Manage Azure DevOps projects",
		Long:  "Manage Azure DevOps projects including list and set operations.",
	}
)

func init() {
	ProjectsCmd.AddCommand(list.ProjectsListCommand)
	ProjectsCmd.AddCommand(set.ProjectSetCommand)
}
