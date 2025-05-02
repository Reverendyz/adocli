package item

import (
	"fmt"

	"github.com/reverendyz/adocli/cmd/item/list"
	"github.com/spf13/cobra"
)

// itemCmd represents the item command
var ItemCmd = &cobra.Command{
	Use:   "item",
	Short: "Displays azure devops items in a given project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("item called")
	},
}

func init() {
	ItemCmd.AddCommand(list.ListCmd)
}
