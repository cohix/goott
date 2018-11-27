package action

import (
	client "github.com/cohix/goott-server/service"
	"github.com/cohix/simplcrypto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Client describes a client for a goott-server
type Client struct {
	goottClient client.GoottClient
	keypair     *simplcrypto.KeyPair
	sessionKey  *simplcrypto.SymKey
}

// CreateGoottClient creates a Client
func CreateGoottClient() (*Client, error) {
	conn, err := grpc.Dial("localhost:3687", grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "failed to Dial")
	}

	client := &Client{
		goottClient: client.NewGoottClient(conn),
	}

	return client, nil
}
