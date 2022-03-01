package email

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
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

func (es *Email) VerificationEmail() {

	//	href := fmt.Sprintf("%s?token=%s", url, token)
	e := email.NewEmail()
	e.From = "Ed Fan <ed0176@cchntek.com>"
	e.To = []string{es.Address}
	e.Subject = "Please verify your email address"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte(fmt.Sprintf(`<a href="%s">Please verify your email address</a>`, es.URI))
	//needs to confirm your email address is valid. Please click the link below to confirm you received this mail.
	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "ed0176@cchntek.com", "tllujunjrqezmtbo", "smtp.gmail.com"))
	if err != nil {
		fmt.Println("email error", err)
	}
}

//SetURI
func (es *Email) SetURI(url, query string) {

	es.URI = url + "?" + query
}

func (es *Email) SetToken(token string) {
	es.Token = token
}
