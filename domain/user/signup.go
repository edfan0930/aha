package user

import (
	"fmt"
	"net/http"
	"unicode"

	"github.com/edfan0930/aha/common/email"
	"github.com/edfan0930/aha/domain/response"

	"github.com/edfan0930/aha/utils"

	"github.com/edfan0930/aha/db"

	"github.com/gin-gonic/gin"
)

type (
	SignupConfirm struct {
		LoginRequest
		ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	}
)

//Signup
func Signup(c *gin.Context) {

	r := NewSignupConfirm()
	if err := c.BindJSON(r); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	//VerifyPassword
	if err := VerifyPassword(r.Password); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	token := utils.GenerUUID()
	if token == "" {
		c.JSON(http.StatusInternalServerError, response.Error("gener uuid error"))
	}

	user := db.NewUser(r.Email).Signup(r.Password, token)
	if err := user.Create(db.MainSession, c); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	email.VerificationEmail([]string{r.Email})

	c.JSON(http.StatusOK, r)
}

//NewSignupConfirm
func NewSignupConfirm() *SignupConfirm {
	return &SignupConfirm{
		LoginRequest: LoginRequest{},
	}
}

//VerifyPassword
func VerifyPassword(password string) error {

	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	var specialCharPresent bool
	const minPassLength = 8
	const maxPassLength = 64
	//	var passLen int
	var err error
	length := len(password)
	if length < minPassLength || length > maxPassLength {
		err = fmt.Errorf("password length must be between %d to %d characters long", minPassLength, maxPassLength)
	}

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
		case unicode.IsUpper(ch):
			uppercasePresent = true
		case unicode.IsLower(ch):
			lowercasePresent = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialCharPresent = true
		case ch == ' ':
			err = fmt.Errorf("/n ; %v", err)
		}
	}

	if !lowercasePresent {
		err = fmt.Errorf("lowercase letter missing /n ; %v", err)
	}
	if !uppercasePresent {
		err = fmt.Errorf("uppercase letter missing /n ; %v", err)
	}
	if !numberPresent {
		err = fmt.Errorf("atleast one numeric character required /n ; %v", err)
	}
	if !specialCharPresent {
		err = fmt.Errorf("special character missing /n ; %v", err)
	}

	return err
}
