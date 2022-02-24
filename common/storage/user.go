package storage

import (
	"errors"
	"net/http"

	"github.com/gorilla/sessions"
)

//OauthSignup
func OauthLogin(w http.ResponseWriter, r *http.Request, email string) error {

	s, _ := UserHandler(r)

	s.Values[StorageKey.Email] = email
	s.Values[StorageKey.Logged] = true

	return s.Save(r, w)
}

//GetEmail
func GetEmail(r *http.Request) (string, error) {

	s, err := UserHandler(r)
	if err != nil {
		return "", err
	}

	l := s.Values[StorageKey.Email]
	email, ok := l.(string)
	if !ok {
		return "", errors.New("assert type error")
	}

	return email, nil
}

//GetLoggedOn
func GetLoggedOn(r *http.Request) (bool, error) {

	s, err := UserHandler(r)
	if err != nil {
		return false, err
	}

	l := s.Values[StorageKey.Logged]
	logged, ok := l.(bool)
	if !ok {
		return false, errors.New("assert type error")
	}

	return logged, nil

}

//Login
func Login(s *sessions.Session, email string) *sessions.Session {

	s.Values[StorageKey.Logged] = true
	s.Values[StorageKey.Email] = email
	return ResetMaxAge(s)
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
