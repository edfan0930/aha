package oauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/oauth2"
)

var UUID string

func init() {
	UUID = generUUID()
}

//Exchange
func Exchange(config *oauth2.Config, ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {

	return config.Exchange(ctx, code, opts...)
}

//AuthURL
func AuthURL(config *oauth2.Config) string {

	return config.AuthCodeURL(UUID)
}

//AuthUUID
func AuthUUID() string {
	return UUID
}

//generUUID ...
func generUUID() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}

	return base64.URLEncoding.EncodeToString(b)
}

func Client()