package oauth2

import (
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
)

const (
	GoogleKey    = "479205503773-r674qa8u7b186hbupe43oimrrc9mrhga.apps.googleusercontent.com"
	GoogleSecret = "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"
	FBKey        = "844823353051244"
	FBSecret     = "83fd5fd6bf47f4f0808fc6109427519d"
)

func SetProvider(url string) {

	goth.UseProviders(
		google.New(GoogleKey, GoogleSecret, GenerURI(url, "google"),
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		),
		facebook.New(FBKey, FBSecret, GenerURI(url, "facebook")),
	)
}

//
func SetStore(key []byte) {

	maxAge := 86400 * 30 // 30 days
	isProd := false      // Set to true when serving over https

	store := sessions.NewCookieStore(key)
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	//gothic.Store = store
}

//GenerURI
func GenerURI(url, path string) string {

	return fmt.Sprintf("%v/%v", url, path)
}
