package get

import (
	"fmt"

	"github.com/reverendyz/adocli/pkg/item/get"
	"github.com/spf13/cobra"
)

var (
	GetWorkItemCmd = &cobra.Command{
		Use:   "get",
		Short: "Gets the WorkItem from a given id",
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			err := get.GetWorkItemById(id)
			if err != nil {
				fmt.Printf("error: %v\n", err.Error())
			}
		},
	}
)
