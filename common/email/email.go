package email

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

var (
	FromAndName string
	From        string
)

type (
	Email struct {
		Callback string `json:"callback"`
		To       string `json:"to"`
		Token    string `json:"token"`
	}
)

func InitEmail(from string) {
	FromAndName = fmt.Sprintf("Ed Fan <%s>", from)
	From = from
}

//NewEmail
func NewEmail(to string) *Email {
	return &Email{
		To: to,
	}
}

//VerificationEmail
func (e *Email) VerificationEmail() {

	ne := email.NewEmail()
	ne.From = FromAndName
	ne.To = []string{e.To}
	ne.Subject = "Please verify your email address"
	ne.Text = []byte("Text Body is, of course, supported!")
	ne.HTML = []byte(fmt.Sprintf(`<a href="%s">Please verify your email address</a>`, e.Callback))

	err := ne.Send("smtp.gmail.com:587", smtp.PlainAuth("", From, "ssaxcoohgkxobroj", "smtp.gmail.com"))
	if err != nil {
		fmt.Println("email error", err)
	}
}

//SetURI
func (e *Email) SetURI(url, query string) {

	e.Callback = url + "?" + query
}

//SetToken
func (es *Email) SetToken(token string) {
	es.Token = token
}
