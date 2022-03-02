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
		Mux         sync.RWMutex
		CookieStore *sessions.CookieStore
		Key         []byte
	}

	Session struct {
		S *sessions.Session
		*store
	}
)

var Store *store

func init() {

	NewStore()
}

//NewCookieStore
func NewStore() {

	key := []byte(GenerSessionID())
	Store = &store{
		CookieStore: sessions.NewCookieStore(key),
		Key:         key,
	}
}

func NewSession(s *sessions.Session) *Session {
	return &Session{
		S:     s,
		store: Store,
	}
}

//GenerSessionID
func GenerSessionID() string {
	return utils.GenerUUID()
}

//UserHandler
func UserHandler(r *http.Request) (*sessions.Session, error) {

	return Store.CookieStore.Get(r, UserStore)
}

//PassSecure pass secure verification
func PassSecure(r *http.Request) *sessions.Session {
	s, _ := UserHandler(r)
	return s
}

//Save
func (s *Session) Save(w http.ResponseWriter, r *http.Request) error {

	return s.S.Save(r, w)
}

//SetDelete
func (s *Session) SetDelete() *Session {
	s.S.Options.MaxAge = -1
	return s
}

//ResetMaxAge
func (s *Session) ResetMaxAge() *Session {
	s.S.Options.MaxAge = MaxAge
	return s
}

func (s *Session) SetValue(key, value string) *Session {
	s.S.Values[key] = value
	return s
}

//GetValue
func (s *Session) GetValue(key string) string {

	v := s.S.Values[key]

	return v.(string)
}
