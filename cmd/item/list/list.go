package list

import (
	"github.com/reverendyz/adocli/pkg/item/list"
	"github.com/spf13/cobra"
)

var ItemListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists work items in a given organization and project",
	RunE: func(cmd *cobra.Command, args []string) error {
		return list.ListAllWorkItemTracking()
	},
}
