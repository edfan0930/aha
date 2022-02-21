package oauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/oauth2"
)

var UUID string

func init() {
	UUID = generUUID()
}

//Exchange
func Exchange(config *oauth2.Config, ctx context.Context, code string, token *oauth2.Token, opts ...oauth2.AuthCodeOption) error {

	token, err := config.Exchange(ctx, code, opts...)
	fmt.Println("is exchange err", err)
	return err
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
		fmt.Println("gener UUID failed: ", err)
		return ""
	}

	return base64.URLEncoding.EncodeToString(b)
}
