package action

import (
	"context"
	"os"

	model "github.com/cohix/goott-server/model"
	"github.com/cohix/simplcrypto"
	"github.com/pkg/errors"
)

const tokenEnvKey = "GOOTT_TOKEN"

// Auth authenticates the client from an env var
func (c *Client) Auth(token string) error {
	if token == "" {
		envExists := false

		token, envExists = os.LookupEnv(tokenEnvKey)
		if envExists == false {
			return errors.New("unable to find auth token in GOOTT_TOKEN or --token")
		}
	}
	keypair, err := simplcrypto.GenerateNewKeyPair()
	if err != nil {
		return errors.Wrap(err, "failed to GenerateKeyPair")
	}

	decodedToken, err := simplcrypto.Base64URLDecode(token)
	if err != nil {
		return errors.Wrap(err, "failed to Base64URLDecode")
	}

	tokenSignature, err := keypair.Sign(decodedToken)
	if err != nil {
		return errors.Wrap(err, "failed to Sign token")
	}

	authReq := &model.AuthRequest{
		TokenSignature: tokenSignature,
		PubKey:         keypair.SerializablePubKey(),
	}

	resp, err := c.goottClient.Auth(context.Background(), authReq)
	if err != nil {
		return errors.Wrap(err, "failed to Auth")
	}

	decSessionKey, err := keypair.Decrypt(resp.EncSessionKey)
	if err != nil {
		return errors.Wrap(err, "failed to Decrypt session key")
	}

	sessionKey, err := simplcrypto.SymKeyFromJSON(decSessionKey)
	if err != nil {
		return errors.Wrap(err, "failed to SymKeyFromJSON")
	}

	c.sessionKey = sessionKey
	c.keypair = keypair

	return nil
}
