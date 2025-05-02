package org

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var OrgCmd = &cobra.Command{
	Use: fmt.Sprintf("%s set org <org name>", os.Args[0]),
}

func init() {

}
