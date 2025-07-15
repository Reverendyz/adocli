package item

import (
	"github.com/reverendyz/adocli/cmd/item/get"
	"github.com/reverendyz/adocli/cmd/item/list"
	"github.com/spf13/cobra"
)

// itemCmd represents the item command
var ItemCmd = &cobra.Command{
	Use:   "item",
	Short: "Manage Azure DevOps work items",
	Long:  "Manage Azure DevOps work items in a given project. Use subcommands to list, get, or modify work items.",
}

func init() {
	ItemCmd.AddCommand(list.ItemListCmd)
	ItemCmd.AddCommand(get.GetWorkItemCmd)
}
