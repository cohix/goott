package model

import (
	"crypto/rand"

	"github.com/cohix/simplcrypto"
)

// App describes the settings for the server
type App struct {
	AuthToken string
	Keyset    *simplcrypto.KeySet
}

// InitializeApp sets up the server
func InitializeApp() *App {
	token := generateAuthToken()

	return &App{
		AuthToken: token,
		Keyset:    &simplcrypto.KeySet{},
	}
}

func generateAuthToken() string {
	bytes := make([]byte, 24)
	rand.Read(bytes)

	return simplcrypto.Base64URLEncode(bytes)
}
