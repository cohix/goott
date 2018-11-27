package service

import (
	simplcrypto "github.com/cohix/simplcrypto"
	log "github.com/cohix/simplog"
	"github.com/pkg/errors"
	context "golang.org/x/net/context"
)

// GetSecretMessage allows the client to get the secret message
func (gs *GoottService) GetSecretMessage(ctx context.Context, req *simplcrypto.Message) (*simplcrypto.Message, error) {
	defer log.LogTrace("GetSecretMessage")()

	sessionKey := gs.App.Keyset.SymKeyWithKID(req.KID)
	if sessionKey == nil {
		return nil, errors.New("failed to fetch SymKeyWithKID; session invalid")
	}

	if _, err := sessionKey.Decrypt(req); err != nil {
		return nil, errors.New("failed to Decrypt, session invalid")
	}

	encMessage, err := sessionKey.Encrypt([]byte(gs.Message))
	if err != nil {
		return nil, errors.New("failed to Encrypt message")
	}

	return encMessage, nil
}

// SetSecretMessage allows the client to get the secret message
func (gs *GoottService) SetSecretMessage(ctx context.Context, req *simplcrypto.Message) (*Empty, error) {
	defer log.LogTrace("SetSecretMessage")()

	sessionKey := gs.App.Keyset.SymKeyWithKID(req.KID)
	if sessionKey == nil {
		return nil, errors.New("failed to fetch SymKeyWithKID; session invalid")
	}

	message, err := sessionKey.Decrypt(req)
	if err != nil {
		return nil, errors.New("failed to Decrypt, session invalid")
	}

	gs.Message = string(message)

	return &Empty{}, nil
}
