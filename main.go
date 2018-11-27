package main

import (
	"os"

	"github.com/cohix/goott/action"
	"github.com/cohix/goott/command"
	log "github.com/cohix/simplog"
	"github.com/pkg/errors"
)

func main() {
	client, err := action.CreateGoottClient()
	if err != nil {
		log.LogError(errors.Wrap(err, "failed to CreateGoottCLient"))
		os.Exit(1)
	}

	command.Execute(client)
}
