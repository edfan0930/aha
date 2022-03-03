package oauth2

import (
	"fmt"

	"github.com/edfan0930/aha/env"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
)

func SetProvider(url string) {

	goth.UseProviders(
		google.New(env.GoogleKey, env.GoogleSecret, GenerURI(url, "google"),
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		),
		facebook.New(env.FBKey, env.FBSecret, GenerURI(url, "facebook")),
	)
}

//
func SetStore(key []byte) {

	maxAge := 86400 * 30 // 30 days
	isProd := false      // Set to true when serving over https

	store := sessions.NewCookieStore(key)
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = false // HttpOnly should always be enabled
	store.Options.Secure = isProd

	//gothic.Store = store
}

//GenerURI
func GenerURI(url, path string) string {

	return fmt.Sprintf("%v/%v", url, path)
}
