package email

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func VerificationEmail(address []string) {
	e := email.NewEmail()
	e.From = "Ed Fan <ed0176@cchntek.com>"
	e.To = []string{"ed0176@cchntek.com"}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	e.Subject = "Please verify your email address"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	//needs to confirm your email address is valid. Please click the link below to confirm you received this mail.
	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "ed0176@cchntek.com", "tllujunjrqezmtbo", "smtp.gmail.com"))
	if err != nil {
		fmt.Println("email error", err)
	}
}
