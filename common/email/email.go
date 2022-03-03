package email

import (
	"fmt"
	"net/smtp"

	"github.com/edfan0930/aha/env"

	"github.com/jordan-wright/email"
)

var (
	from string = fmt.Sprintf("Ed Fan <%s>", env.Email)
)

type (
	Email struct {
		URI     string `json:"uri"`
		Address string `json:"address"`
		Token   string `json:"token"`
	}
)

//NewEmail
func NewEmail(address string) *Email {
	return &Email{
		Address: address,
	}
}

//VerificationEmail
func (es *Email) VerificationEmail() {

	e := email.NewEmail()
	e.From = from
	e.To = []string{es.Address}
	e.Subject = "Please verify your email address"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte(fmt.Sprintf(`<a href="%s">Please verify your email address</a>`, es.URI))

	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", env.Email, env.EmailPassword, "smtp.gmail.com"))
	if err != nil {
		fmt.Println("email error", err)
	}
}

//SetURI
func (es *Email) SetURI(url, query string) {

	es.URI = url + "?" + query
}

//SetToken
func (es *Email) SetToken(token string) {
	es.Token = token
}
