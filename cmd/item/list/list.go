package list

import (
	"fmt"
	"os"

	"github.com/reverendyz/adocli/pkg/item/list"
	"github.com/spf13/cobra"
)

var ItemListCmd = &cobra.Command{
	Use:   "list",
	Short: "lists workitems in a given org and project",
	Run: func(cmd *cobra.Command, args []string) {
		err := list.ListAllWorkItemTracking()
		if err != nil {
			fmt.Printf("error: %v\n", err.Error())
			os.Exit(1)
		}
	},
}
