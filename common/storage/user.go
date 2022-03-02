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

//GetVerified
func (s *Session) GetVerified(r *http.Request) string {

	return s.GetValue(StorageKey.Verified)
}

//Login
func (s *Session) Login(w http.ResponseWriter, r *http.Request, email, name string, verified bool) error {

	var v string
	if verified {

		v = "true"
	}
	if !verified {

		v = "false"
	}

	s.SetValue(StorageKey.Email, email)
	s.SetValue(StorageKey.Logged, "true")
	s.SetValue(StorageKey.Name, name)
	s.SetValue(StorageKey.Verified, v)
	s.ResetMaxAge()

	return s.Save(w, r)
}

//Verified
func (s *Session) Verified(w http.ResponseWriter, r *http.Request) error {

	s.SetValue(StorageKey.Verified, "true")
	return s.Save(w, r)
}

//Logout
func (s *Session) Logout(w http.ResponseWriter, r *http.Request) error {

	s.SetDelete()

	return s.Save(w, r)
}
