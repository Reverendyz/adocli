package get

import (
	"github.com/reverendyz/adocli/pkg/item/get"
	"github.com/spf13/cobra"
)

var (
	GetWorkItemCmd = &cobra.Command{
		Use:   "get <work-item-id>",
		Short: "Gets the WorkItem from a given id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return get.GetWorkItemById(args[0])
		},
	}
)
