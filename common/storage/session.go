package storage

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

func Get() {

	fmt.Println("store", Store)
}

var Store = sessions.NewCookieStore([]byte(GenerSessionID()))

//GenerSessionID
func GenerSessionID() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	fmt.Println("uid", b)
	return base64.URLEncoding.EncodeToString(b)
}

func Handler(w http.ResponseWriter, r *http.Request, key string) {
	Store.Get(r, key)
}

//SetDelete
func SetDelete(s *sessions.Session) {
	s.Options.MaxAge = -1
}
