package set

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var SetCmd = &cobra.Command{
	Use: fmt.Sprintf("%s set [org]", os.Args[0]),
}

func init() {
	SetCmd.AddCommand(org.OrgCmd)
}
