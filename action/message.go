package action

import (
	"context"

	"github.com/pkg/errors"
)

// SetSecretMessage sets the secret message on the server
func (c *Client) SetSecretMessage(message string) error {
	encMessage, err := c.sessionKey.Encrypt([]byte(message))
	if err != nil {
		return errors.Wrap(err, "failed to Encrypt")
	}

	_, err = c.goottClient.SetSecretMessage(context.Background(), encMessage)

	return err
}

// GetSecretMessage sets the secret message on the server
func (c *Client) GetSecretMessage() (string, error) {
	encAuthMessage, err := c.sessionKey.Encrypt([]byte("randomnonsensegoeshere"))
	if err != nil {
		return "", errors.Wrap(err, "failed to Encrypt")
	}

	encMessage, err := c.goottClient.GetSecretMessage(context.Background(), encAuthMessage)
	if err != nil {
		return "", errors.Wrap(err, "failed to GetSecretMessage")
	}

	decMessage, err := c.sessionKey.Decrypt(encMessage)
	if err != nil {
		return "", errors.Wrap(err, "failed to Decrypt message")
	}

	return string(decMessage), err
}
