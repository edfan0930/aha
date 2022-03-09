package email

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

var (
	fromAndName string
	from        string
	callbackURL string
)

type (
	Email struct {
		Callback      string `json:"callback"`
		To            string `json:"to"`
		Token         string `json:"token"`
		CallBackQuery string
	}
)

//InitEmail init email field
func InitEmail(sender, url string) {
	fromAndName = fmt.Sprintf("Ed Fan <%s>", from)
	from = sender
	callbackURL = url
}

//NewEmail
func NewEmail(to string) *Email {
	return &Email{
		To: to,
	}
}

//VerificationEmail send email
func (e *Email) VerificationEmail() error {

	ne := email.NewEmail()
	ne.From = fromAndName
	ne.To = []string{e.To}
	ne.Subject = "Please verify your email address"
	ne.Text = []byte("Text Body is, of course, supported!")
	ne.HTML = []byte(fmt.Sprintf(`<a href="%s">Please verify your email address</a>`, e.Callback))

	err := ne.Send("smtp.gmail.com:587", smtp.PlainAuth("", from, "ssaxcoohgkxobroj", "smtp.gmail.com"))
	if err != nil {
		fmt.Println("email error", err)

	}
	return err
}

//SetURI
func (e *Email) SetURI() *Email {

	e.Callback = callbackURL + "?" + e.CallBackQuery
	return e
}

//SetToken
func (e *Email) SetToken(token string) {

	e.Token = token
}

//SetQuery set callback query
func (e *Email) SetQuery(token, account string) *Email {

	e.CallBackQuery = fmt.Sprintf("token=%s&account=%s", token, account)
	return e
}
