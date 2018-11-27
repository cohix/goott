package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/cohix/goott/action"
	log "github.com/cohix/simplog"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func setMessageCmd(client *action.Client) *cobra.Command {
	var format string

	cmd := &cobra.Command{
		Use:   "message",
		Short: "set the secret message",
		Long:  `set the secret message`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := client.Auth(authToken); err != nil {
				log.LogError(errors.Wrap(err, "failed to Auth"))
				os.Exit(1)
			}

			message := args[0]

			if format == "caps" {
				message = strings.ToUpper(message)
			} else if format == "underscores" {
				message = strings.Replace(message, " ", "_", -1)
			}

			log.LogInfo(fmt.Sprintf("setting message: %s", message))

			if err := client.SetSecretMessage(message); err != nil {
				log.LogError(errors.Wrap(err, "failed to SetSecretMessage"))
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVar(&format, "format", "", "either 'caps' or 'underscores' to set message in all caps or with underscores instead of spaces")

	return cmd
}

func getMessageCmd(client *action.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "message",
		Short: "get the secret message",
		Long:  `get the secret message`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := client.Auth(authToken); err != nil {
				log.LogError(errors.Wrap(err, "failed to Auth"))
				os.Exit(1)
			}

			message, err := client.GetSecretMessage()
			if err != nil {
				log.LogError(errors.Wrap(err, "failed to GetSecretMessage"))
				os.Exit(1)
			}

			fmt.Println(message)
		},
	}

	return cmd
}
