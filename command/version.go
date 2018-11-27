package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

const goottVersion = 0.1

func versionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "list the version of goott",
		Long:  `list the version of goott`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(goottVersion)
		},
	}

	return cmd
}
