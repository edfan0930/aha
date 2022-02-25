package storage

import (
	"net/http"
)

//OauthSignup
func (s *Session) OauthLogin(w http.ResponseWriter, r *http.Request, email, name string) error {

	s.SetValue(StorageKey.Email, email)
	s.SetValue(StorageKey.Logged, "true")
	s.SetValue(StorageKey.Name, name)
	s.ResetMaxAge()
	return s.Save(w, r)
}

//GetEmail
func (s *Session) GetEmail(r *http.Request) string {

	return s.GetValue(StorageKey.Email)
}

//GetLoggedOn
func (s *Session) GetLoggedOn(r *http.Request) string {

	return s.GetValue(StorageKey.Logged)
}

//GetName
func (s *Session) GetName(r *http.Request) string {

	return s.GetValue(StorageKey.Name)
}

//Login
func (s *Session) Login(w http.ResponseWriter, r *http.Request, email, name string) error {

	s.SetValue(StorageKey.Email, email)
	s.SetValue(StorageKey.Logged, "true")
	s.SetValue(StorageKey.Name, name)
	s.ResetMaxAge()

	return s.Save(w, r)
}

//Logout
func (s *Session) Logout(w http.ResponseWriter, r *http.Request) error {

	s.SetDelete()

	return s.Save(w, r)
}
