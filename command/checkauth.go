package command

import (
	"os"

	"github.com/cohix/goott/action"
	log "github.com/cohix/simplog"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func checkAuthCmd(client *action.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "ensure that the present auth token is valid",
		Long:  `verify the auth token provided by the GOOTT_TOKEN env var or --token arg flag`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := client.Auth(authToken); err != nil {
				log.LogError(errors.Wrap(err, "failed to Auth"))
				os.Exit(1)
			}

			log.LogInfo("auth is valid")
		},
	}

	return cmd
}
