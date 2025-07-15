package item

import (
	"github.com/reverendyz/adocli/cmd/item/get"
	"github.com/reverendyz/adocli/cmd/item/list"
	"github.com/spf13/cobra"
)

// itemCmd represents the item command
var ItemCmd = &cobra.Command{
	Use:   "item",
	Short: "Displays azure devops items in a given project",
}

func init() {
	ItemCmd.AddCommand(list.ItemListCmd)
	ItemCmd.AddCommand(get.GetWorkItemCmd)
}
