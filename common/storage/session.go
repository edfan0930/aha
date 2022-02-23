package storage

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/sessions"
)

type (
	store struct {
		mux     sync.RWMutex
		Session *sessions.CookieStore
		Key     []byte
	}
)

var Store *store

func init() {

	NewCookieStore()
}

//NewCookieStore
func NewCookieStore() {

	key := []byte(GenerSessionID())
	Store = &store{
		Session: sessions.NewCookieStore(key),
		Key:     key,
	}
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

func GetEmail(s *sessions.Session) (string, error) {

	l := s.Values[StorageKey.Email]
	email, ok := l.(string)
	if !ok {
		return "", errors.New("assert type error")
	}

	return email, nil
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

//UserHandler
func UserHandler(w http.ResponseWriter, r *http.Request) (*sessions.Session, error) {

	return Store.Session.Get(r, "user")
}

//SetDelete
func SetDelete(s *sessions.Session) {
	s.Options.MaxAge = -1
}
