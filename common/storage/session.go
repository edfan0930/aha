package storage

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

func init() {

	Store = sessions.NewCookieStore([]byte(GenerSessionID()))
}

//GenerSessionID
func GenerSessionID() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("gener session ID failed:", err)
		return ""
	}

	return base64.URLEncoding.EncodeToString(b)
}

func Get() {

	fmt.Println("store", Store)
}

func Login(s *sessions.Session) {
	s.Values[StorageKey.Logged] = true
}

//LoggedOn get logged bool
func LoggedOn(s *sessions.Session) (bool, error) {
	l := s.Values[StorageKey.Logged]

	logged, ok := l.(bool)
	if !ok {
		return false, errors.New("assert type error")
	}

	return logged, nil
}

func Handler(w http.ResponseWriter, r *http.Request, key string) {
	Store.Get(r, key)
}

//SetDelete
func SetDelete(s *sessions.Session) {
	s.Options.MaxAge = -1
}
