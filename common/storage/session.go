package storage

import (
	"net/http"
	"sync"

	"github.com/edfan0930/aha/utils"

	"github.com/gorilla/sessions"
)

const (
	UserStore = "user"
	MaxAge    = 86400 * 30
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
	return utils.GenerUUID()
}

//Save
func Save(s *sessions.Session, w http.ResponseWriter, r *http.Request) error {

	return s.Save(r, w)
}

//UserHandler
func UserHandler(r *http.Request) (*sessions.Session, error) {

	return Store.Session.Get(r, UserStore)
}

//SetDelete
func SetDelete(s *sessions.Session) {
	s.Options.MaxAge = -1
}

//ResetMaxAge
func ResetMaxAge(s *sessions.Session) *sessions.Session {
	s.Options.MaxAge = MaxAge
	return s
}
