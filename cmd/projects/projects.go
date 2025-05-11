package projects

import (
	"github.com/reverendyz/adocli/cmd/projects/list"
	"github.com/spf13/cobra"
)

var ProjectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "handles projects commands",
}

func init() {
	ProjectsCmd.AddCommand(list.ProjectsListCommand)
}
