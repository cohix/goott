package command

import (
	"fmt"
	"os"

	"github.com/cohix/goott/action"
	log "github.com/cohix/simplog"
	"github.com/spf13/cobra"
)

var authToken string

// Execute runs the tool
func Execute(client *action.Client) {
	cmd := arrangeCommands(client)

	if err := cmd.Execute(); err != nil {
		log.LogError(err)
		os.Exit(1)
	}
}

func arrangeCommands(client *action.Client) *cobra.Command {
	root := rootCmd()
	root.AddCommand(versionCmd())

	checkCmd := checkCmd()
	checkCmd.AddCommand(checkAuthCmd(client))
	root.AddCommand(checkCmd)

	getCmd := getCmd()
	getCmd.AddCommand(getMessageCmd(client))
	root.AddCommand(getCmd)

	setCmd := setCmd()
	setCmd.AddCommand(setMessageCmd(client))
	root.AddCommand(setCmd)

	return root
}

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goot",
		Short: "goot is an example command-line client for a gRPC service",
		Long:  `An example of building command-line clients in Go, and a primer for gRPC`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s\n\nuse `goott --help` to view available commands, and `goott version` to list version\n", cmd.Short)
		},
	}

	cmd.PersistentFlags().StringVar(&authToken, "token", "", "--token overrides the GOOTT_TOKEN env var")

	return cmd
}
